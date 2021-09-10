package com.wgtwo.api.common;

public enum Environment {
    SANDBOX("sandbox.api.wgtwo.com", 443),
    PRODUCTION("api.wgtwo.com", 443);

    public final String host;
    public final int port;
    public final String target;

    Environment(String host, int port) {
        this.host = host;
        this.port = port;
        this.target = host + ":" + port;
    }
}
