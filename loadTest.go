package main

func loadTest(users int, endpoint string) string {

	for i := 0; i < users; i++ {
		go httpCall(endpoint)
	}
	return "Success"
}
