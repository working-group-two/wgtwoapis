package com.wgtwo.api.util.auth;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

public class Clients {
    public enum Environment {
        DEV("apigateway.dub.dev.wgtwo.com"),
        PROD("api.wgtwo.com");

        private final String host;

        Environment(String host) {
            this.host = host;
        }
    }

    public static ManagedChannel createChannel(Environment environment) {
        return ManagedChannelBuilder.forAddress(environment.host, 443).build();
    }
}
