package com.wgtwo.api.util.auth;

import io.grpc.CallCredentials;
import io.grpc.Metadata;
import io.grpc.Status;
import java.nio.charset.StandardCharsets;
import java.util.Base64;
import java.util.concurrent.Executor;

public class OperatorToken extends CallCredentials {
  private static final Metadata.Key<String> AUTH_HEADER =
      Metadata.Key.of("authorization", Metadata.ASCII_STRING_MARSHALLER);

  private String clientId;
  private String clientSecret;

  public OperatorToken(String clientId, String clientSecret) {
    this.clientId = clientId;
    this.clientSecret = clientSecret;
  }

  @Override
  public void applyRequestMetadata(
      RequestInfo requestInfo, Executor executor, MetadataApplier metadataApplier) {
    executor.execute(
        () -> {
          try {
            Metadata headers = new Metadata();
            headers.put(AUTH_HEADER, "Basic " + base64());
            metadataApplier.apply(headers);
          } catch (Throwable e) {
            metadataApplier.fail(Status.UNAUTHENTICATED.withCause(e));
          }
        });
  }

  @Override
  public void thisUsesUnstableApi() {}

  private String base64() {
    String credentials = clientId + ":" + clientSecret;
    return Base64.getEncoder().encodeToString(credentials.getBytes(StandardCharsets.US_ASCII));
  }
}
