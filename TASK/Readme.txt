Programming Task Objective
--------------------------

Write a simple web application in the Go programming language that performs
JSON-to-HTML conversion.

The application shall be a standard HTTP server which takes POST requests,
reads the body in the request, converts it from JSON to HTML according to
a predefined template and returns the HTML as the body of the response.

The executable shall accept the following command-line arguments:

  1. `-p <PortNumber>` specifying the port number on which the HTTP
      server will listen.

  2. `-t <TemplateFilePath>` specifying the HTML template file.

Upon startup, the executable shall parse the HTML template file using the
standard `html/template` library.

The HTTP server shall accept incoming requests in parallel.

A `GET` request to the root URL should display a simple web page with a form and
a submit button. The form will contain a `<textarea>` for inputting the JSON.
Submit action will be a `POST` request to the `/render` URL.

Each `POST` request to the `/render` URL shall be processed.

Valid JSON has the following format:

    {
      "threatName": "Win32/Rbot",
      "category": "trojan",
      "size": 437289,
      "detectionDate": "2019-04-01",
      "variants": [
        {
          "name": "Win32/TrojanProxy.Emotet.A",
          "dateAdded": "2019-04-10"
        },
        {
          "name": "Win32/TrojanProxy.Emotet.B",
          "dateAdded": "2019-04-22"
        },
        ...
      ]
    }

An example template might look like: (see also `threat.html.tmpl` file)

    ...
    <h2>threatName {{.ThreatName}}</h2>
    <h2>category {{.Category}}</h2>
    ...

An HTML main page body might look like:

    <!DOCTYPE html>
    <html>
    <body>
    <form action="/render" method="POST">
    <label for="json_input">JSON input:</label><br/>
    <textarea rows="4" cols="50" id="json_input"
              name="json_input">
    </textarea><br/>
    <input type="submit" value="Submit"/>
    </form>
    </body>
    </html>

If any of the keys is missing, default value will be used (0, false, empty
string). Any extra unknown keys will be ignored. The HTML template must be
able to display every key given in the above example.

The resulting data shall be returned as HTTP response, with the `text/html;
charset=UTF-8` content type. Any other request, as well as malformed requests,
shall generate an error response with a reasonable status code and message.

Use standard Golang libraries, where appropriate:

  - <https://golang.org/pkg/flag/> (for command-line parameter parsing)
  - <https://golang.org/pkg/net/http/> (for HTTP server)
  - <https://golang.org/pkg/encoding/json/> (for JSON parsing)
  - <https://golang.org/pkg/html/template/> (for HTML template processing)
  - any other standard library

Note that solving the assignment doesn't require any use of javascript.

Programming Task Evaluation
---------------------------

When evaluating your solution we will consider:

  1. Correctness.
  2. Code clarity.
  3. Efficiency.
  4. Documentation.
  5. Robustness.