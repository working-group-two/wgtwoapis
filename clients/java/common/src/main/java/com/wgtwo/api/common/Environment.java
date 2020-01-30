package com.wgtwo.api.common;

public enum Environment {
    DEV("apigateway.dub.dev.wgtwo.com"),
    PROD("api.wgtwo.com");

    public final String host;

    Environment(String host) {
        this.host = host;
    }
}
