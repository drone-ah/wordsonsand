---
title: "Microservices vs Monolith: Real World Tradeoffs"
date: 2024-07-17T09:48:25+01:00
categories: architecture
tags:
  - architecture
  - java
  - java-spring
  - reactive-web
---

When starting a new backend system for a contract I was on, one of the early
decisions I had to make was whether to lean into a monolith or adopt a
microservices approach. While common wisdom offers strong opinions on both ends
of the spectrum, in reality, the choice often hinges on organizational
constraints as much as on technical purity.

### Reactive vs Traditional Spring Web

I began by reviewing
[performance comparisons](https://filia-aleks.medium.com/microservice-performance-battle-spring-mvc-vs-webflux-80d39fd81bf0)
between Spring MVC and WebFlux. Reactive Web generally comes out ahead in
benchmarks, but that doesn’t tell the whole story.

In our use case—web notifications—the benefit of reactive patterns depends
heavily on how data is delivered. If we were polling, the advantage would be
limited. However, with Server-Sent Events (SSE), Spring’s support aligns
directly with Reactive Web, making WebFlux the more appropriate choice for this
part of the system.

### The Deployment Constraint

Ideally, I would have started with a monolith: a single deployable artifact
combining both the Kafka Streams logic and the API. This option would have
simplified initial development and allowed us to iterate quickly. But at the
client, the platform does not allow deploying a Kafka Streams app and an API
within the same Kubernetes deployment.

This effectively rules out a true monolith, even for a prototype.

### Options Considered

#### Shared Library with Thin Deployments

A middle ground was to build the core logic in a shared library and have
lightweight deployments wrap around it. This would allow the streams app and the
API to share code without needing to make HTTP calls between them.

The downside: these services are no longer independently deployable. But given
our team size and velocity goals, this compromise might be acceptable.

#### Full Microservices

Another option was to separate the services entirely:

- **Streams service** (Kafka, plus domain-specific logic)
- **Web API** (for delivering notifications)
- **Subscription API** (managing notification subscriptions)

This adheres more closely to the single responsibility principle, especially as
we move from PoC to MVP. However, it adds deployment and coordination overhead.

#### Application Profiles

A third hacky option was to control which parts of the app run using
environment-based profiles. For example, we could disable Kafka in dev or use
conditional beans to keep deployments clean. While not ideal long-term, it
offers flexibility for early stages.

### Conclusion

Constraints matter. While I lean toward monoliths for rapid delivery in small
teams, platform limitations forced a hybrid approach. We intend to evolve into
microservices over time, but only when the benefits clearly outweigh the cost.

Have you faced similar deployment constraints that shaped your architecture? I'd
love to hear how you navigated them.
