package main

import (
	"context"
	"os"
	// "encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"

	"database/sql"

	fdk "github.com/fnproject/fdk-go"
	_ "github.com/go-sql-driver/mysql"
)

const (
	Host     = "DB_HOST"
	User     = "DB_USER"
	Password = "DB_PASSWORD"
	Db       = "DB_DB"
)

var errorHtml = `<!DOCTYPE html>
<html>

<head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="shortcut icon" type="image/x-icon" href="/resources/images/favicon.png">
    <title>PetClinic :: a Spring Framework demonstration</title>

    <!--[if lt IE 9]>
      <script src="https://oss.maxcdn.com/html5shiv/3.7.2/html5shiv.min.js"></script>
      <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
      <![endif]-->

    <link rel="stylesheet" href="/resources/css/petclinic.css" />
</head>

<body>
    <nav class="navbar navbar-default" role="navigation">
        <div class="container">
            <div class="navbar-header">
                <a class="navbar-brand" href="/">
                    <span></span>
                </a>
                <button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#main-navbar">
                    <span class="sr-only">
                        <os-p>Toggle navigation</os-p>
                    </span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                </button>
            </div>
            <div class="navbar-collapse collapse" id="main-navbar">
                <ul class="nav navbar-nav navbar-right">

                    <li>
                        <a href="">
                            <span class="glyphicon  glyphicon-null" aria-hidden="true"></span>
                            <span></span>
                        </a>
                    </li>

                    <li>
                        <a href="/" title="home page">
                            <span class="glyphicon  glyphicon-home" aria-hidden="true"></span>
                            <span>Home</span>
                        </a>
                    </li>

                    <li>
                        <a href="/owners/find" title="find owners">
                            <span class="glyphicon  glyphicon-search" aria-hidden="true"></span>
                            <span>Find owners</span>
                        </a>
                    </li>

                    <li class="active">
                        <a href="/vets.html" title="veterinarians">
                            <span class="glyphicon  glyphicon-th-list" aria-hidden="true"></span>
                            <span>Veterinarians</span>
                        </a>
                    </li>

                    <li>
                        <a href="/contact.html" title="contact">
                            <span class="glyphicon  glyphicon-envelope" aria-hidden="true"></span>
                            <span>Contact</span>
                        </a>
                    </li>

                    <li>
                        <a href="/oups" title="trigger a RuntimeException to see how it is handled">
                            <span class="glyphicon  glyphicon-warning-sign" aria-hidden="true"></span>
                            <span>Error</span>
                        </a>
                    </li>

                </ul>
            </div>
        </div>
    </nav>
    <div class="container-fluid">
        <div class="container xd-container">
            <h2>New Vet</h2>
            <div class="alert alert-danger" role="alert">
				There was an error when adding a new vet.
				%v
            </div>

            <br/>
            <div class="container">
                <div class="row">
                    <div class="col-12 text-center">
                        Modified by gloo&trade;
                    </div>
                </div>
            </div>
        </div>
    </div>

    <script src="/webjars/jquery/2.2.4/jquery.min.js"></script>
    <script src="/webjars/jquery-ui/1.11.4/jquery-ui.min.js"></script>
    <script src="/webjars/bootstrap/3.3.6/js/bootstrap.min.js"></script>

</body>

</html>
`

var successHtml = `<!DOCTYPE html>
<html>

<head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="shortcut icon" type="image/x-icon" href="/resources/images/favicon.png">
    <title>PetClinic :: a Spring Framework demonstration</title>

    <!--[if lt IE 9]>
      <script src="https://oss.maxcdn.com/html5shiv/3.7.2/html5shiv.min.js"></script>
      <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
      <![endif]-->

    <link rel="stylesheet" href="/resources/css/petclinic.css" />
</head>

<body>
    <nav class="navbar navbar-default" role="navigation">
        <div class="container">
            <div class="navbar-header">
                <a class="navbar-brand" href="/">
                    <span></span>
                </a>
                <button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#main-navbar">
                    <span class="sr-only">
                        <os-p>Toggle navigation</os-p>
                    </span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                </button>
            </div>
            <div class="navbar-collapse collapse" id="main-navbar">
                <ul class="nav navbar-nav navbar-right">

                    <li>
                        <a href="">
                            <span class="glyphicon  glyphicon-null" aria-hidden="true"></span>
                            <span></span>
                        </a>
                    </li>

                    <li>
                        <a href="/" title="home page">
                            <span class="glyphicon  glyphicon-home" aria-hidden="true"></span>
                            <span>Home</span>
                        </a>
                    </li>

                    <li>
                        <a href="/owners/find" title="find owners">
                            <span class="glyphicon  glyphicon-search" aria-hidden="true"></span>
                            <span>Find owners</span>
                        </a>
                    </li>

                    <li class="active">
                        <a href="/vets.html" title="veterinarians">
                            <span class="glyphicon  glyphicon-th-list" aria-hidden="true"></span>
                            <span>Veterinarians</span>
                        </a>
                    </li>

                    <li>
                        <a href="/contact.html" title="contact">
                            <span class="glyphicon  glyphicon-envelope" aria-hidden="true"></span>
                            <span>Contact</span>
                        </a>
                    </li>

                    <li>
                        <a href="/oups" title="trigger a RuntimeException to see how it is handled">
                            <span class="glyphicon  glyphicon-warning-sign" aria-hidden="true"></span>
                            <span>Error</span>
                        </a>
                    </li>

                </ul>
            </div>
        </div>
    </nav>
    <div class="container-fluid">
        <div class="container xd-container">
            <h2>New Vet</h2>
            <div class="alert alert-success" role="alert">
                New vet successfully added.
            </div>

            <br/>
            <div class="container">
                <div class="row">
                    <div class="col-12 text-center">
                        Modified by gloo&trade;
                    </div>
                </div>
            </div>
        </div>
    </div>

    <script src="/webjars/jquery/2.2.4/jquery.min.js"></script>
    <script src="/webjars/jquery-ui/1.11.4/jquery-ui.min.js"></script>
    <script src="/webjars/bootstrap/3.3.6/js/bootstrap.min.js"></script>

</body>

</html>
`
var settings = map[string]string{
	User:     "root",
	Password: "demo",
	Host:     "petclinic-db.default.svc.cluster.local",
	Db:       "petclinic",
}

func main() {
	for k := range settings {
		v := os.Getenv(k)
		if v != "" {
			settings[k] = v
		}
	}
	fdk.Handle(fdk.HandlerFunc(myHandler))
}

type Vet struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	City      string `json:"city"`
}

func addVet(v Vet) error {

	connstring := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", settings[User], settings[Password], settings[Host], settings[Db])
	db, err := sql.Open("mysql", connstring)

	if err != nil {
		return err
	}
	_, err = db.Query("INSERT INTO vets (first_name, last_name, city) VALUES (?,?,?)", v.FirstName, v.LastName, v.City)
	if err != nil {
		return err
	}
	return nil
	// connect and say hi
}

func myHandler(_ context.Context, in io.Reader, out io.Writer) {
	b, err := ioutil.ReadAll(in)

	if err != nil {
		// fmt.Fprintf(out, "don't be a child")
		fmt.Fprintf(out, errorHtml, err)
		return
	}
	vs, err := url.ParseQuery(string(b))

	if err != nil {
		// fmt.Fprintf(out, "still on't be a child")
		fmt.Fprintf(out, errorHtml, err)
		return
	}
	vet := Vet{
		FirstName: vs.Get("firstName"),
		LastName:  vs.Get("lastName"),
		City:      vs.Get("city"),
	}

	err = addVet(vet)

	if err != nil {
		// fmt.Fprintf(out, "%v", err)
		fmt.Fprintf(out, errorHtml, err)
		return
	}
	fmt.Fprintf(out, successHtml)

}
