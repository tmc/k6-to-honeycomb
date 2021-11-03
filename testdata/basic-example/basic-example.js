// Basic example to demonstrate functionality.

import {get,post} from 'k6/http';

export default function() {
	const r1 = get("http://localhost:9090/get");
	const r2 = post("http://localhost:9090/post", {});
}
