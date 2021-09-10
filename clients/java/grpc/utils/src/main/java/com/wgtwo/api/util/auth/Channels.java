package com.wgtwo.api.util.auth;

import com.wgtwo.api.common.Environment;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

public class Channels {
    public static ManagedChannel createChannel(Environment environment) {
        return ManagedChannelBuilder.forAddress(environment.host, 443).build();
    }

    public static ManagedChannelBuilder<?> createChannelBuilder(Environment environment) {
        return ManagedChannelBuilder.forAddress(environment.host, 443);
    }
}
