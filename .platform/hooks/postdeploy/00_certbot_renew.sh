#!/bin/bash
# Renew if the cert is <30 days from expiry, then reload nginx
certbot renew --quiet --deploy-hook "systemctl reload nginx"