## API_Call

A simple Golang application to make API calls and retrieve data using free APIs from [cricketdata.org](https://cricketdata.org/).

* PLEASE DON'T GET HARDCODE API CALL ever, this is for test purpose only. 

## Learnings

* Created a structure that mirrors the data type. In this case, my data type was structured, and was able to create a corresponding struct for it on Golang. 
* Used `http.Get(url)` to make an API call.
* Read the response from the API call using `io.ReadAll(resp.Body)`.
* Decoded the JSON output using `json.Unmarshal`.


