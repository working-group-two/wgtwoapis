package com.wgtwo.api.auth;

import com.wgtwo.api.common.OperatorToken;
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

    public static Token createToken() {
        return new Token();
    }

    public static class Token {
        public OperatorToken operatorToken(String clientId, String clientSecret) {
            return new OperatorToken(clientId, clientSecret);
        }

        public UserToken userToken(String token) {
            return new UserToken(token);
        }
    }
}
