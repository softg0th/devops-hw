import http from 'k6/http';
import { check } from 'k6';

export let options = {
    vus: 100,
    duration: '100s',
};


export default function () {
    const url = 'http://127.0.0.1:9111/api/animals';
    const payload = JSON.stringify({
        "Name" :   "a",
        "Type":    "string",
        "Color":   "string",
        "StoreID": 1,
        "Age":     1,
        "Price":  1
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    let res = http.post(url, payload, params);

    check(res, {
        'status is 201': (r) => r.status === 201,
    });
}