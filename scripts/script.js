import http from 'k6/http';
import { sleep, check } from 'k6';

export default function () {
    const res = http.get(__ENV.TARGET_URL);
    sleep(1);

    const checkRes = check(res, {
        'status is 200': (r) => r.status === 200
    });
}
