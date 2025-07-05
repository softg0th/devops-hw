import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
  vus: 100,
  duration: '3m',
};

export default function () {
  const url = 'http://127.0.0.1:9111/api/stores';
  const payload = JSON.stringify({
    Name: 'eee',
    Address: 'hello world',
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const res = http.post(url, payload, params);

  check(res, {
    'status is 200': (r) => r.status === 200,
  });

  sleep(0.1);
}