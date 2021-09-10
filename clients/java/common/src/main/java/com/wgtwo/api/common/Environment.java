package com.wgtwo.api.common;

public enum Environment {
    SANDBOX("sandbox.api.wgtwo.com"),
    PRODUCTION("api.wgtwo.com");

    public final String host;

    Environment(String host) {
        this.host = host;
    }
}
