# Application Platform Management API

- [Application Platform Management API](#application-platform-management-api)
  - [Functional Requirements](#functional-requirements)
    - [Tenant application onboarding APIs](#tenant-application-onboarding-apis)
    - [User onboarding APIs](#user-onboarding-apis)
    - [File Uploader APIs](#file-uploader-apis)
    - [User verification APIs](#user-verification-apis)
  - [Observability Requirements](#observability-requirements)
  - [Tech Stack](#tech-stack)

> [!NOTE]
> STATE: `early-development`
> At the moment, this project is more to explore Golang than build a product.

Multi-tenant application managment API to view and manage applications and users

## Functional Requirements

### Tenant application onboarding APIs

- [ ] Onboard tenant application and user sign in preferences.
- [ ] Get all onboarded tenant applications.

### User onboarding APIs

- [ ] Register users to applications using email, password.
- [ ] Register users to applications using oauth (Google, GitHub)
- [ ] User sign in to generate token for onboarded applications.

### File Uploader APIs

- [ ] Upload user profile image in required resolutions.
- [ ] Retrieve user profile image in specified resolution.

### User verification APIs

Used for non oauth clients

- [ ] Send & verify user email verification
- [ ] Send & verify user phone verification

## Observability Requirements

I have few options, but I'll have to explore each options for setup:

- ELK
- Dynatrace
- Grafana

## Tech Stack

- [x] _Golang_: why not?
- [x] _PostgreSQL_: relation database
- [ ] _AWS S3_: Object storage for storing images etc.
- [ ] _Kafka_/_RabbitMQ_: Process messages for asynchronous API flows.
- [ ] _DigitalOcean_: Deployment to VPS.
