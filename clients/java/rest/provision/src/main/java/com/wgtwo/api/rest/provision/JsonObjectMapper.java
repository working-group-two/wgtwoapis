package com.wgtwo.api.rest.provision;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.JsonDeserializationContext;
import com.google.gson.JsonDeserializer;
import com.google.gson.JsonElement;
import com.google.gson.JsonParseException;
import com.google.gson.JsonPrimitive;
import com.google.gson.JsonSerializationContext;
import com.google.gson.JsonSerializer;
import java.lang.reflect.Type;
import java.time.Instant;
import kong.unirest.GenericType;
import kong.unirest.ObjectMapper;

public class JsonObjectMapper implements ObjectMapper {
    private final Gson gson = new GsonBuilder().registerTypeAdapter(Instant.class, new InstantTypeConverter()).create();

    public <T> T readValue(String value, Class<T> valueType) {
        return this.gson.fromJson(value, valueType);
    }

    public <T> T readValue(String value, GenericType<T> genericType) {
        return this.gson.fromJson(value, genericType.getType());
    }

    public String writeValue(Object value) {
        return this.gson.toJson(value);
    }

    private static class InstantTypeConverter
            implements JsonSerializer<Instant>, JsonDeserializer<Instant> {
        @Override
        public JsonElement serialize(Instant src, Type srcType, JsonSerializationContext context) {
            return new JsonPrimitive(src.toEpochMilli());
        }

        @Override
        public Instant deserialize(JsonElement json, Type type, JsonDeserializationContext context)
                throws JsonParseException {
            return Instant.parse(json.getAsString());
        }
    }
}

