package gateway;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.netflix.zuul.ZuulFilter;
import com.netflix.zuul.context.RequestContext;
import com.netflix.zuul.exception.ZuulException;
import com.netflix.zuul.http.ServletInputStreamWrapper;
import lombok.extern.slf4j.Slf4j;
import org.apache.commons.lang.StringEscapeUtils;
import org.apache.commons.lang3.StringUtils;
import org.springframework.cloud.context.config.annotation.RefreshScope;
import org.springframework.cloud.netflix.zuul.filters.support.FilterConstants;
import org.springframework.stereotype.Component;
import org.springframework.util.StreamUtils;

import javax.servlet.ServletInputStream;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletRequestWrapper;
import java.io.IOException;
import java.io.InputStream;
import java.nio.charset.Charset;
import java.util.*;

@Component
@Slf4j
@RefreshScope
public class XSSFilter extends ZuulFilter {
    @Override
    public String filterType() {
        return FilterConstants.PRE_TYPE;
    }

    @Override
    public int filterOrder() {
        return 1;
    }

    @Override
    public boolean shouldFilter() {
        return true;
    }

    @Override
    public Object run() throws ZuulException {
        RequestContext requestContext = RequestContext.getCurrentContext();
        HttpServletRequest request = requestContext.getRequest();
        String contentType = request.getContentType();
        if (StringUtils.isBlank(contentType)) {
            return null;
        } else if (StringUtils.equals(contentType, "application/x-www-form-urlencoded")
                || StringUtils.equals(contentType, "application/x-www-form-urlencoded;charset=UTF-8")) {
            Map<String, String[]> parameterMap = request.getParameterMap();
            Iterator it_d = parameterMap.entrySet().iterator();
            while (it_d.hasNext()) {
                Map.Entry<String, String[]> entry_d = (Map.Entry) it_d.next();
                String key = entry_d.getKey();
                String[] value = entry_d.getValue();
                if (value != null) {
                    List<String> strings = Arrays.asList(value);
                    for (int i = 0; i < strings.size(); i++) {
                        strings.set(i, StringEscapeUtils.escapeHtml(strings.get(i)));
                        strings.set(i, StringEscapeUtils.escapeJavaScript(strings.get(i)));
                    }
                }
                parameterMap.put(key, value);
            }
            String newBody = JSON.toJSONString(parameterMap);
            final byte[] reqBodyBytes = newBody.getBytes();
            requestContext.setRequest(new HttpServletRequestWrapper(request) {

                @Override
                public ServletInputStream getInputStream() {
                    return new ServletInputStreamWrapper(reqBodyBytes);
                }

                @Override
                public int getContentLength() {
                    return reqBodyBytes.length;
                }

                @Override
                public long getContentLengthLong() {
                    return reqBodyBytes.length;
                }
            });

        } else if (StringUtils.equals(contentType, "application/json")
                || StringUtils.equals(contentType, "application/json;charset=UTF-8")) {
            try {
                InputStream in = requestContext.getRequest().getInputStream();
                String body = StreamUtils.copyToString(in, Charset.forName("UTF-8"));

                JSONObject json = JSON.parseObject(body);
                Map<String, Object> map = json;
                Map<String, Object> mapJson = new HashMap<>();
                for (Map.Entry<String, Object> entry : map.entrySet()) {
                    mapJson.put(entry.getKey(), cleanXSS(entry.getValue().toString()));
                }
                String newBody = JSON.toJSONString(mapJson);
                final byte[] reqBodyBytes = newBody.getBytes();
                requestContext.setRequest(new HttpServletRequestWrapper(request) {

                    @Override
                    public ServletInputStream getInputStream() {
                        return new ServletInputStreamWrapper(reqBodyBytes);
                    }

                    @Override
                    public int getContentLength() {
                        return reqBodyBytes.length;
                    }

                    @Override
                    public long getContentLengthLong() {
                        return reqBodyBytes.length;
                    }
                });
            } catch (IOException e) {
                log.error("xss filter read parameter abnormal", e);
            }
        }
        try {
            InputStream in = requestContext.getRequest().getInputStream();
            String body = StreamUtils.copyToString(in, Charset.forName("UTF-8"));
            System.out.println(body);
        } catch (
                IOException e) {
            log.error("xss filter read parameter abnormal", e);
        }
        return null;
    }

    private String cleanXSS(String value) {
        if (StringUtils.isBlank(value)) {
            return value;
        }
        System.out.println("HENLO");
        value = StringEscapeUtils.escapeHtml(value);
        value = StringEscapeUtils.escapeJavaScript(value);
        value = value.replaceAll("\\\\", "");
        return value;
    }
}