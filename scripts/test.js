import http from 'k6/http'

export let options = {
    vus: 5,
    duration: '5s',
    // iteration: 10
}

export default function() {
    http.get('http://host.docker.internal:8000/products')
}