# Nginx crash course

> web server which can also be used as a reverse proxy, load balancer, mail 
> proxy and HTTP cache. 

In this short guide, we will cover how one can leverage nginx as a web server,
briefly cover load balancing, and discuss security.

## Installing nginx
LOL!

```Dockerfile
FROM nginx:1.15.5                                                               
                                                                                
COPY nginx.conf /etc/nginx/nginx.conf                                           
                                                                                
                                                                                
CMD ["nginx", "-g", "daemon off;"]
```


## Check your configuration is valid
```
nginx -t
```

## Semantics

Nginx has modules.
Here we will cover only the http module. 
As such, you will see this a lot:
```
http {
    ...
}
```

The overall format for nginx's configuration is:
```
<section> {
    <directive> <params>;
}
```

This way we can define different configuration contexts to pull a service or
services together into a coherent api.

#### Useful global configuration directives

* `error_log` you can specify the level - `debug`, `info`, `notice`, `warn`,    
`error`, `crit`, `alert`, and `emerg`.

* `worker_process` number of worker processes that will be started. For
  cpu-bound workloads usually setting it to 1x number of cores will give will
  good performance. For IO-bound workloads, maybe 1.5x or 2x the number of
  cores.

* `worker_connections` maximum number of open connections a worker process may
  have open at any given time.


#### Common http directives
* [`listen`](http://nginx.org/en/docs/http/ngx_http_core_module.html#listen)
  listen for request coming from `<address[:port]>`.

* [`location`](http://nginx.org/en/docs/http/ngx_http_core_module.html#location)
  defines how to handle requests.


## References

### Documentation
- [Alphabetical index of variables](http://nginx.org/en/docs/varindex.html) 
    For things like `$request_uri`, `$host`, `$args`.

- [Understanding Nginx Server and Location Block Selection Algorithms](https://www.digitalocean.com/community/tutorials/understanding-nginx-server-and-location-block-selection-algorithms)
    Match the request uri to a service.

### HTTPS
- [Configuring HTTPS servers](http://nginx.org/en/docs/http/configuring_https_servers.html)

- [Strong ciphers](https://cipherli.st/)

- [Strong SSL security on nginx](https://raymii.org/s/tutorials/Strong_SSL_Security_On_nginx.html)

- [HTTP Strict Transport Security (HSTS) and NGINX](https://www.nginx.com/blog/http-strict-transport-security-hsts-and-nginx/)

### Load balancing
- [Using nginx as HTTP load balancer](http://nginx.org/en/docs/http/load_balancing.html)


### HTTP cache
- [A Guide to Caching with NGINX and NGINX Plus](https://www.nginx.com/blog/nginx-caching-guide/)
