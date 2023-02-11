import axios from 'axios';

const apiV1 = axios.create({
  baseURL: 'http://localhost/api/v1',  
});

export default apiV1;