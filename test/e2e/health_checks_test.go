package e2e_test

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gogo/protobuf/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/api/v2/cluster"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/api/v2/core"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"github.com/solo-io/gloo/projects/gloo/pkg/translator"
	"github.com/solo-io/gloo/test/services"
	"github.com/solo-io/gloo/test/v1helpers"
	glootest "github.com/solo-io/gloo/test/v1helpers/test_grpc_service/glootest/protos"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

var _ = Describe("Health Checks", func() {
	var (
		ctx            context.Context
		cancel         context.CancelFunc
		testClients    services.TestClients
		envoyInstance  *services.EnvoyInstance
		tu             *v1helpers.TestUpstream
		writeNamespace string
	)

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())
		defaults.HttpPort = services.NextBindPort()
		defaults.HttpsPort = services.NextBindPort()

		var err error
		envoyInstance, err = envoyFactory.NewEnvoyInstance()
		Expect(err).NotTo(HaveOccurred())

		writeNamespace = defaults.GlooSystem
		ro := &services.RunOptions{
			NsToWrite: writeNamespace,
			NsToWatch: []string{"default", writeNamespace},
			WhatToRun: services.What{
				DisableGateway: false,
				DisableUds:     true,
				// test relies on FDS to discover the grpc spec via reflection
				DisableFds: false,
			},
		}
		testClients = services.RunGlooGatewayUdsFds(ctx, ro)
		err = envoyInstance.RunWithRole(writeNamespace+"~gateway-proxy-v2", testClients.GlooPort)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		if envoyInstance != nil {
			_ = envoyInstance.Clean()
		}
		cancel()
	})

	basicReq := func(b []byte) func() (string, error) {
		return func() (string, error) {
			// send a request with a body
			var buf bytes.Buffer
			buf.Write(b)
			res, err := http.Post(fmt.Sprintf("http://%s:%d/test", "localhost", defaults.HttpPort), "application/json", &buf)
			if err != nil {
				return "", err
			}
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			return string(body), err
		}
	}

	Context("regression for config", func() {

		BeforeEach(func() {

			tu = v1helpers.NewTestGRPCUpstream(ctx, envoyInstance.LocalAddr(), 1)
			_, err := testClients.UpstreamClient.Write(tu.Upstream, clients.WriteOpts{})
			Expect(err).NotTo(HaveOccurred())

		})

		tests := []struct {
			Name  string
			Check *core.HealthCheck
		}{
			{
				Name: "http",
				Check: &core.HealthCheck{
					HealthChecker: &core.HealthCheck_HttpHealthCheck_{
						HttpHealthCheck: &core.HealthCheck_HttpHealthCheck{
							Path: "xyz",
						},
					},
				},
			},
			{
				Name: "tcp",
				Check: &core.HealthCheck{
					HealthChecker: &core.HealthCheck_TcpHealthCheck_{
						TcpHealthCheck: &core.HealthCheck_TcpHealthCheck{
							Send: &core.HealthCheck_Payload{
								Payload: &core.HealthCheck_Payload_Text{
									Text: "AAAA",
								},
							},
							Receive: []*core.HealthCheck_Payload{
								{
									Payload: &core.HealthCheck_Payload_Text{
										Text: "AAAA",
									},
								},
							},
						},
					},
				},
			},
		}

		for _, v := range tests {
			v := v
			It(v.Name, func() {
				us, err := testClients.UpstreamClient.Read(tu.Upstream.Metadata.Namespace, tu.Upstream.Metadata.Name, clients.ReadOpts{})
				Expect(err).NotTo(HaveOccurred())
				v.Check.Timeout =&translator.DefaultHealthCheckTimeout
				v.Check.Interval = &translator.DefaultHealthCheckInterval
				v.Check.HealthyThreshold = translator.DefaultGogoThreshold
				v.Check.UnhealthyThreshold = translator.DefaultGogoThreshold
				us.GetUpstreamSpec().HealthChecks = []*core.HealthCheck{v.Check}

				_, err = testClients.UpstreamClient.Write(us, clients.WriteOpts{
					OverwriteExisting: true,
				})
				Expect(err).NotTo(HaveOccurred())

				vs := getGrpcVs(writeNamespace, tu.Upstream.Metadata.Ref())
				_, err = testClients.VirtualServiceClient.Write(vs, clients.WriteOpts{})
				Expect(err).NotTo(HaveOccurred())

				body := []byte(`{"str": "foo"}`)

				testRequest := basicReq(body)

				Eventually(testRequest, 30, 1).Should(Equal(`{"str":"foo"}`))

				Eventually(tu.C).Should(Receive(PointTo(MatchFields(IgnoreExtras, Fields{
					"GRPCRequest": PointTo(Equal(glootest.TestRequest{Str: "foo"})),
				}))))
			})
		}

		It("outlier detection", func() {
			us, err := testClients.UpstreamClient.Read(tu.Upstream.Metadata.Namespace, tu.Upstream.Metadata.Name, clients.ReadOpts{})
			Expect(err).NotTo(HaveOccurred())
			us.GetUpstreamSpec().OutlierDetection = &cluster.OutlierDetection{
				Interval: &types.Duration{Seconds: 1},
			}

			_, err = testClients.UpstreamClient.Write(us, clients.WriteOpts{
				OverwriteExisting: true,
			})
			Expect(err).NotTo(HaveOccurred())

			vs := getGrpcVs(writeNamespace, tu.Upstream.Metadata.Ref())
			_, err = testClients.VirtualServiceClient.Write(vs, clients.WriteOpts{})
			Expect(err).NotTo(HaveOccurred())

			body := []byte(`{"str": "foo"}`)

			testRequest := basicReq(body)

			Eventually(testRequest, 30, 1).Should(Equal(`{"str":"foo"}`))

			Eventually(tu.C).Should(Receive(PointTo(MatchFields(IgnoreExtras, Fields{
				"GRPCRequest": PointTo(Equal(glootest.TestRequest{Str: "foo"})),
			}))))
		})
	})

	Context("e2e + GRPC", func() {

		BeforeEach(func() {
			tu = v1helpers.NewTestGRPCUpstream(ctx, envoyInstance.LocalAddr(), 5)
			_, err := testClients.UpstreamClient.Write(tu.Upstream, clients.WriteOpts{})
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() error { return envoyInstance.SetPanicThreshold() }, time.Second*5, time.Second/4).Should(BeNil())

			tu = v1helpers.NewTestGRPCUpstream(ctx, envoyInstance.LocalAddr(), 5)
			_, err = testClients.UpstreamClient.Write(tu.Upstream, clients.WriteOpts{})
			Expect(err).NotTo(HaveOccurred())

			us, err := testClients.UpstreamClient.Read(tu.Upstream.Metadata.Namespace, tu.Upstream.Metadata.Name, clients.ReadOpts{})
			Expect(err).NotTo(HaveOccurred())

			us.GetUpstreamSpec().HealthChecks = []*core.HealthCheck{
				{
					Timeout:            &translator.DefaultHealthCheckTimeout,
					Interval:           &translator.DefaultHealthCheckInterval,
					UnhealthyThreshold: translator.DefaultGogoThreshold,
					HealthyThreshold:   translator.DefaultGogoThreshold,
					HealthChecker: &core.HealthCheck_GrpcHealthCheck_{
						GrpcHealthCheck: &core.HealthCheck_GrpcHealthCheck{
							ServiceName: "TestService",
						},
					},
				},
			}

			_, err = testClients.UpstreamClient.Write(us, clients.WriteOpts{
				OverwriteExisting: true,
			})
			Expect(err).NotTo(HaveOccurred())

			vs := getGrpcVs(writeNamespace, tu.Upstream.Metadata.Ref())
			_, err = testClients.VirtualServiceClient.Write(vs, clients.WriteOpts{})
			Expect(err).NotTo(HaveOccurred())
		})

		It("Fail all but one GRPC health check", func() {
			liveService := tu.FailGrpcHealthCheck()
			body := []byte(`{"str": "foo"}`)
			testRequest := basicReq(body)

			numRequests := 5

			for i := 0; i < numRequests; i++ {
				Eventually(testRequest, 30, 1).Should(Equal(`{"str":"foo"}`))
			}

			for i := 0; i < numRequests; i++ {
				select {
				case v := <-tu.C:
					Expect(v.Port).To(Equal(liveService.Port))
				case <-time.After(5 * time.Second):
					Fail("channel did not receive proper response in time")
				}
			}
		})
	})

})
