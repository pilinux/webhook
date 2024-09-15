# webhook

processing webhooks easy

## Supported platforms

- [x] [Resend](https://resend.com/docs/dashboard/webhooks/introduction)
  - Implemented event types
    - [x] `email.sent`
    - [x] `email.delivered`
    - [x] `email.delivery_delayed`
    - [x] `email.complained`
    - [x] `email.bounced`
    - [x] `email.opened`
    - [x] `email.clicked`
    - [x] `contact.created`
    - [x] `contact.updated`
    - [x] `contact.deleted`

- [x] [stripe events](https://docs.stripe.com/webhooks)
  - Go SDK version: `v79`
  - API version: `2024-06-20`
  - [Handle webhook versioning](https://docs.stripe.com/webhooks/versioning)
  - [CLI](https://docs.stripe.com/stripe-cli)
    - login: `stripe login`
    - listen: `stripe listen --latest --skip-verify --forward-to localhost:4242/stripe_webhooks`
    - trigger: `stripe trigger <event>`, e.g. `stripe trigger customer.subscription.created`
  - Implemented [event types](https://docs.stripe.com/api/events/types)
    - [ ] `account.application.authorized`
    - [ ] `account.application.deauthorized`
    - [ ] `account.external_account.created`
    - [ ] `account.external_account.deleted`
    - [ ] `account.external_account.updated`
    - [ ] `account.updated`
    - [ ] `application_fee.created`
    - [ ] `application_fee.refund.updated`
    - [ ] `application_fee.refunded`
    - [ ] `balance.available`
    - [ ] `billing_portal.configuration.created`
    - [ ] `billing_portal.configuration.updated`
    - [ ] `billing_portal.session.created`
    - [ ] `billing.alert.triggered`
    - [ ] `capability.updated`
    - [ ] `cash_balance.funds_available`
    - [x] [charge](https://docs.stripe.com/api/charges)
      - [x] `charge.captured`
      - [x] `charge.dispute.closed`
      - [x] `charge.dispute.created`
      - [x] `charge.dispute.funds_reinstated`
      - [x] `charge.dispute.funds_withdrawn`
      - [x] `charge.dispute.updated`
      - [x] `charge.expired`
      - [x] `charge.failed`
      - [x] `charge.pending`
      - [x] `charge.refund.updated`
      - [x] `charge.refunded`
      - [x] `charge.succeeded`
      - [x] `charge.updated`
    - [ ] `checkout.session.async_payment_failed`
    - [ ] `checkout.session.async_payment_succeeded`
    - [ ] `checkout.session.completed`
    - [ ] `checkout.session.expired`
    - [ ] `climate.order.canceled`
    - [ ] `climate.order.created`
    - [ ] `climate.order.delayed`
    - [ ] `climate.order.delivered`
    - [ ] `climate.order.product_substituted`
    - [ ] `climate.product.created`
    - [ ] `climate.product.pricing_updated`
    - [ ] `coupon.created`
    - [ ] `coupon.deleted`
    - [ ] `coupon.updated`
    - [ ] `credit_note.created`
    - [ ] `credit_note.updated`
    - [ ] `credit_note.voided`
    - [ ] `customer_cash_balance_transaction.created`
    - [ ] `customer.created`
    - [ ] `customer.updated`
    - [ ] `customer.deleted`
    - [ ] `customer.discount.created`
    - [ ] `customer.discount.deleted`
    - [ ] `customer.discount.updated`
    - [ ] `customer.source.created`
    - [ ] `customer.source.deleted`
    - [ ] `customer.source.expiring`
    - [ ] `customer.source.updated`
    - [x] [customer.subscription](https://docs.stripe.com/api/subscriptions)
      - [x] `customer.subscription.created`
      - [x] `customer.subscription.deleted`
      - [x] `customer.subscription.paused`
      - [x] `customer.subscription.pending_update_applied`
      - [x] `customer.subscription.pending_update_expired`
      - [x] `customer.subscription.resumed`
      - [x] `customer.subscription.trial_will_end`
      - [x] `customer.subscription.updated`
    - [ ] `customer.tax_id.created`
    - [ ] `customer.tax_id.deleted`
    - [ ] `customer.tax_id.updated`
    - [ ] `entitlements.active_entitlement_summary.updated`
    - [ ] `file.created`
    - [ ] `financial_connections.account.created`
    - [ ] `financial_connections.account.deactivated`
    - [ ] `financial_connections.account.disconnected`
    - [ ] `financial_connections.account.reactivated`
    - [ ] `financial_connections.account.refreshed_balance`
    - [ ] `financial_connections.account.refreshed_ownership`
    - [ ] `financial_connections.account.refreshed_transactions`
    - [ ] `identity.verification_session.canceled`
    - [ ] `identity.verification_session.created`
    - [ ] `identity.verification_session.processing`
    - [ ] `identity.verification_session.redacted`
    - [ ] `identity.verification_session.requires_input`
    - [ ] `identity.verification_session.verified`
    - [x] [invoice](https://docs.stripe.com/api/invoices)
      - [x] `invoice.created`
      - [x] `invoice.deleted`
      - [x] `invoice.finalization_failed`
      - [x] `invoice.finalized`
      - [x] `invoice.marked_uncollectible`
      - [x] `invoice.overdue`
      - [x] `invoice.paid`
      - [x] `invoice.payment_action_required`
      - [x] `invoice.payment_failed`
      - [x] `invoice.payment_succeeded`
      - [x] `invoice.sent`
      - [x] `invoice.upcoming`
      - [x] `invoice.updated`
      - [x] `invoice.voided`
      - [x] `invoice.will_be_due`
    - [x] [invoiceitem](https://docs.stripe.com/api/invoiceitems)
      - [x] `invoiceitem.created`
      - [x] `invoiceitem.deleted`
    - [ ] `issuing_authorization.created`
    - [ ] `issuing_authorization.request`
    - [ ] `issuing_authorization.updated`
    - [ ] `issuing_card.created`
    - [ ] `issuing_card.updated`
    - [ ] `issuing_cardholder.created`
    - [ ] `issuing_cardholder.updated`
    - [ ] `issuing_dispute.closed`
    - [ ] `issuing_dispute.created`
    - [ ] `issuing_dispute.funds_reinstated`
    - [ ] `issuing_dispute.funds_rescinded`
    - [ ] `issuing_dispute.submitted`
    - [ ] `issuing_dispute.updated`
    - [ ] `issuing_personalization_design.activated`
    - [ ] `issuing_personalization_design.deactivated`
    - [ ] `issuing_personalization_design.rejected`
    - [ ] `issuing_personalization_design.updated`
    - [ ] `issuing_token.created`
    - [ ] `issuing_token.updated`
    - [ ] `issuing_transaction.created`
    - [ ] `issuing_transaction.updated`
    - [ ] `mandate.updated`
    - [x] [payment_intent](https://docs.stripe.com/api/payment_intents)
      - [x] `payment_intent.amount_capturable_updated`
      - [x] `payment_intent.canceled`
      - [x] `payment_intent.created`
      - [x] `payment_intent.partially_funded`
      - [x] `payment_intent.payment_failed`
      - [x] `payment_intent.processing`
      - [x] `payment_intent.requires_action`
      - [x] `payment_intent.succeeded`
    - [ ] `payment_link.created`
    - [ ] `payment_link.updated`
    - [ ] `payment_method.attached`
    - [ ] `payment_method.automatically_updated`
    - [ ] `payment_method.detached`
    - [ ] `payment_method.updated`
    - [ ] `payout.canceled`
    - [ ] `payout.created`
    - [ ] `payout.failed`
    - [ ] `payout.paid`
    - [ ] `payout.reconciliation_completed`
    - [ ] `payout.updated`
    - [ ] `person.created`
    - [ ] `person.deleted`
    - [ ] `person.updated`
    - [ ] `plan.created`
    - [ ] `plan.deleted`
    - [ ] `plan.updated`
    - [ ] `price.created`
    - [ ] `price.deleted`
    - [ ] `price.updated`
    - [ ] `product.created`
    - [ ] `product.deleted`
    - [ ] `product.updated`
    - [ ] `promotion_code.created`
    - [ ] `promotion_code.updated`
    - [ ] `quote.accepted`
    - [ ] `quote.canceled`
    - [ ] `quote.created`
    - [ ] `quote.finalized`
    - [ ] `quote.will_expire`
    - [ ] `radar.early_fraud_warning.created`
    - [ ] `radar.early_fraud_warning.updated`
    - [ ] `refund.created`
    - [ ] `refund.updated`
    - [ ] `reporting.report_run.failed`
    - [ ] `reporting.report_run.succeeded`
    - [ ] `reporting.report_type.updated`
    - [ ] `review.closed`
    - [ ] `review.opened`
    - [x] [setup_intent](https://docs.stripe.com/api/setup_intents)
      - [x] `setup_intent.canceled`
      - [x] `setup_intent.created`
      - [x] `setup_intent.requires_action`
      - [x] `setup_intent.setup_failed`
      - [x] `setup_intent.succeeded`
    - [ ] `sigma.scheduled_query_run.created`
    - [ ] `source.canceled`
    - [ ] `source.chargeable`
    - [ ] `source.failed`
    - [ ] `source.mandate_notification`
    - [ ] `source.refund_attributes_required`
    - [ ] `source.transaction.created`
    - [ ] `source.transaction.updated`
    - [ ] `subscription_schedule.aborted`
    - [ ] `subscription_schedule.canceled`
    - [ ] `subscription_schedule.completed`
    - [ ] `subscription_schedule.created`
    - [ ] `subscription_schedule.expiring`
    - [ ] `subscription_schedule.released`
    - [ ] `subscription_schedule.updated`
    - [ ] `tax_rate.created`
    - [ ] `tax_rate.updated`
    - [ ] `tax.settings.updated`
    - [ ] `terminal.reader.action_failed`
    - [ ] `terminal.reader.action_succeeded`
    - [ ] `test_helpers.test_clock.advancing`
    - [ ] `test_helpers.test_clock.created`
    - [ ] `test_helpers.test_clock.deleted`
    - [ ] `test_helpers.test_clock.internal_failure`
    - [ ] `test_helpers.test_clock.ready`
    - [ ] [topup](https://docs.stripe.com/api/topups)
      - [ ] `topup.canceled`
      - [ ] `topup.created`
      - [ ] `topup.failed`
      - [ ] `topup.reversed`
      - [ ] `topup.succeeded`
    - [ ] `transfer.created`
    - [ ] `transfer.reversed`
    - [ ] `transfer.updated`
