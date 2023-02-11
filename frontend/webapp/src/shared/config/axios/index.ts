import axios from 'axios';

const apiV1 = axios.create({
  baseURL: 'http://127.0.0.1:4000/api/v1',  
});

export default apiV1;