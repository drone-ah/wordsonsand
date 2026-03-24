---
title: "Even Light Gets Heavier"
date: 2026-03-24T10:56:05Z
categories:
  - leadership
tags:
  - leadership
  - ways-of-working
  - production
  - decision-making
  - ai
---

A dedicated input type is better than reusing your domain model at the API
boundary. Test layers matter. Writing log statements as you go saves the poor
soul (probably you) debugging blind at 10pm. You know all of this.

This isn't about any of that.

It's about the fact that none of those decisions show up in the metrics that
matter to the people making hiring and delivery calls. The cost is immediate and
visible. The return is delayed, quiet, and arrives in the form of things that
didn't happen — the investigation that took two hours instead of two days, the
API change that didn't bleed into the domain model, the bug that the structure
caught before it shipped.

Sprint velocity captures the extra day. It doesn't capture what that day bought.

This is not a new problem. Most engineers who've been around long enough have
felt it from both sides - made the careful call and got measured on the
slowness, or inherited the codebase built entirely for speed and paid the tax.
The measurement system was already broken. It has been rewarding the appearance
of velocity over the thing velocity is supposed to serve.

This was true long before anyone was generating code with AI. The PR process in
a lot of teams was already largely theatrical — review comments on naming
conventions while the architectural decisions slipped through unquestioned,
approvals given because the diff was too large to meaningfully read. The gate
was already not doing much. We brushed it under the carpet and moved on.

AI tooling is changing the volume of code moving through that process by an
order of magnitude. The pressure to remove the gate entirely — to trust the
output, to ship faster - is only growing. The faster-is-better incentive that
was already making review ineffective is about to be handed a much larger
surface to work on.

Many years ago, I pitched full redevlopment of a ticketing system from a PHP
based system to a Java EE system because it was struggling to scale.

It probably needed a couple of years to build. They wanted it in six months. I
accepted the challenge.

We built and deployed the system in eight months. We spent the next year fixing
it.

The client then rebuilt it in-house.

When AI runs this experiment at scale, who takes it back?
