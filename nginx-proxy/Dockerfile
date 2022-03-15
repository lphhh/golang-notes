FROM jwilder/nginx-proxy
RUN { \
      echo 'server_tokens off;'; \
      echo 'client_max_body_size 1024m;'; \
      echo 'proxy_buffer_size 256k;'; \
      echo 'proxy_buffers 4 256k;'; \
      echo 'proxy_busy_buffers_size 256k;'; \
      echo 'proxy_temp_file_write_size 256k;'; \
    } > /etc/nginx/conf.d/my_proxy.conf