import http from 'k6/http';
import { check } from 'k6';

export default function () {
  // Define the base URL of your API
  const baseUrl = 'http://localhost:3000'; 

  // Define the endpoint you want to test
  const endpoint = '/beef/summary';

  // Make a GET request to the API
  const response = http.get(`${baseUrl}${endpoint}`);

  // Check the response status code
  check(response, {
    'Status is 200': (r) => r.status === 200,
  });

 
}
