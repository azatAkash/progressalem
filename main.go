package main

func main() {
	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoiYWFrYXNoIiwiZXhwIjoxNzI4NDAzMzk3fQ.32G5pFVb3Zwk-yEeCHhjPj_fX5sDuBd1OvE6DokGsnQ"

	htmlContent := extractData(accessToken)

	printExtracted(htmlContent)
}
