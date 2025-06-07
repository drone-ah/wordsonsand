---
layout: post
title:
  Using CXF Interceptors to do some magic around your web service calls [1105]
date: 2011-11-06 12:24:28.000000000 +00:00
type: post
parent_id: "0"
published: true
password: ""
status: publish
categories:
  - Java (EE)
tags:
  - Apache CXF
  - Exception handling
  - software engineering
  - Spring Framework
  - web services
meta:
  _publicize_pending: "1"
  _edit_last: "48492462"
  oc_commit_id: http://drone-ah.com/2011/11/06/using-cxf-interceptors-to-do-some-magic-around-your-web-service-calls-1105/1320582271
  restapi_import_id: 591d994f7aad5
  original_post_id: "715"
  _wp_old_slug: "715"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:52"
permalink: "/2011/11/06/using-cxf-interceptors-to-do-some-magic-around-your-web-service-calls-1105/"
---

We use JBossWS CXF for a heavily utilised enterprise system. It links into
spring to pick up and execute beans. We have a bunch of exceptions that could
get thrown.

To simplify it, the code was originally written to create an anonymous class a
la Runnable which is wrapped around a try catch block. The exceptions that are
thrown are then converted to a soap fault and passed back.

```java
private SOAPFaultException convertToSoapException(ApplicationException e)
{
    try {
        if(null == soapFactory) {
            soapFactory = SOAPFactory.newInstance();
        }
        SOAPFault sf = soapFactory.createFault();
        sf.setFaultString( e.getMessage() );
        sf.setFaultCode( Integer.toString(e.getErrorCode()) );
        return new SOAPFaultException( sf );
    } catch(SOAPException soapException) {
        throw new RuntimeException( soapException );
    }
}
```

Nothing inherently wrong with this. However, there are a couple of issues with
this in that each soap method is set to throw an _ApplicationException_ and
there is not further documentation of which of the subclasses are actually
relevant to that method.

In a runtime environment, this is not hugely relevant. However, when generating
documentation from the WSDL's, it is.

To resolve this, we changed each method to throw their relevant exception, and
wrote an interceptor to pick up the exception and convert it...

The first step was to write an interceptor which is surprisingly simple
and straightforward.

```java
public class ExampleSoapFaultInterceptor extends AbstractSoapInterceptor {

    Logger log = Logger.getLogger(getClass());

    public QuarkSoapFaultInterceptor() {
        super(Phase.MARSHAL);
    }

    @Override
    public void handleMessage(SoapMessage message) throws Fault {
        Fault f = (Fault) message.getContent(Exception.class);

        Throwable cause = f.getCause();
        if (cause instanceof ApplicationException) {
            log.info("Exception Thrown", cause);
            QuarkException e = (QuarkException) cause;
            f.setFaultCode(new QName("", String.valueOf(e.getErrorCode())));

        } else {
            log.warn("Unexpected Exception thrown ", cause);
        }

    }
}
```

The class doesn't need to extend the AbstractSoapInterceptor but you would then
have to manually implement a number of mechanisms that the abstract class
provides.

The constructor simply defines where this interceptor should be inserted into as
part of the chain. There is a whole bunch of different places where it can be
inserted based on this. Information can be
[found in their documentation.](http://cxf.apache.org/docs/interceptors.html "Apache CXF Interceptors")

To insert this interceptor into a web service, the service needs to be annotated
as follows:

```java
@WebService(endpointInterface = "uk.co.kraya.example.WebService")
@OutFaultInterceptors(interceptors= {"uk.co.kraya.example.interceptors.ExampleSoapFaultInterceptor"})
public class ExampleWebService implements ExampleAPI {
```

Each exception thrown from the web service will now be intercepted. You can then
include any further further information in the soap fault.

In this particular case, the only thing that gets done is to update the fault
code with the error code set against the exception. Any additional information
can be set against the soap fault at this point. You can also log as we are
doing.
