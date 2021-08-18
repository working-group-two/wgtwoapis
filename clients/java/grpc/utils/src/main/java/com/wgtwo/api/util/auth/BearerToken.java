package com.wgtwo.api.util.auth;

import io.grpc.CallCredentials;
import io.grpc.Metadata;
import io.grpc.Status;
import java.util.concurrent.Executor;
import java.util.function.Supplier;

public class BearerToken extends CallCredentials {
  private static final Metadata.Key<String> AUTH_HEADER =
      Metadata.Key.of("authorization", Metadata.ASCII_STRING_MARSHALLER);

  private Supplier<String> tokenSupplier;

  public BearerToken(Supplier<String> tokenSupplier) {
    this.tokenSupplier = tokenSupplier;
  }

  @Override
  public void applyRequestMetadata(
      RequestInfo requestInfo, Executor executor, MetadataApplier metadataApplier) {
    executor.execute(
        () -> {
          try {
            Metadata headers = new Metadata();
            String token = tokenSupplier.get();
            headers.put(AUTH_HEADER, "Bearer " + token);
            metadataApplier.apply(headers);
          } catch (Throwable e) {
            metadataApplier.fail(Status.UNAUTHENTICATED.withCause(e));
          }
        });
  }

  @Override
  public void thisUsesUnstableApi() {}
}
