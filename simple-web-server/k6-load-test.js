import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
  stages: [
    { duration: '10s', target: 100 },
    { duration: '10s', target: 200 },
    { duration: '10s', target: 300 },
    { duration: '30s', target: 600 },
  ],
};

export default function () {
  let res = http.get('http://localhost:8080/weather');
  check(res, { 'status was 200': (r) => r.status == 200 });
  sleep(0.5);
}