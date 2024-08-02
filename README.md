Instructions to run

Build and Run the Docker Container

Open PowerShell and navigate to the directory containing build.ps1, then run:


.\build.ps1


Endpoints

Capture Visitor: POST /capture
Request Body: {"url": "http://foo.com", "visitor": "Alice"}
Query Visitors: GET /query
Open the HTML Page

Open index.html in a browser. The page will fetch and display the visitor data from the API.

