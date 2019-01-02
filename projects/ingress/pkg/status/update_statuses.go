package status

// runningAddresses returns a list of IP addresses and/or FQDN where the
// ingress controller is currently running
func (s *statusSyncer) runningAddresses(ingressSvcName, ingressSvcNamespace string) ([]string, error) {
	var addrs []string

	svc, err := s.Client.CoreV1().Services(ns).Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	if svc.Spec.Type == apiv1.ServiceTypeExternalName {
		addrs = append(addrs, svc.Spec.ExternalName)
		return addrs, nil
	}

	for _, ip := range svc.Status.LoadBalancer.Ingress {
		if ip.IP == "" {
			addrs = append(addrs, ip.Hostname)
		} else {
			addrs = append(addrs, ip.IP)
		}
	}

	addrs = append(addrs, svc.Spec.ExternalIPs...)
	return addrs, nil

}
