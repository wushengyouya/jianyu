tag-api-test:
	curl -X POST http://127.0.0.1:8000/api/v1/tags \
		-H 'Content-Type: application/json' \
		-H 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfa2V5IjoiZWRkeWNqeSIsImFwcF9zZWNyZXQiOiJnby1wcm9ncmFtbWluZy10b3VyLWJvb2siLCJpc3MiOiJibG9nLXNlcnZpY2UiLCJleHAiOjE3NDEyNjMzNTV9.QU5Q7HVmzxYyLqVDfJaduwWP45GJS4y0Gv197XPaIhw' \
		-d '{"name":"GO学习111", "created_by":"eddycjy"}'
get-token:
	curl -X POST 'http://127.0.0.1:8000/auth' -F 'app_key=eddycjy' -F 'app_secret=go-programming-tour-book'

