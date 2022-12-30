#!/bin/bash

psql dev -U ubuntu -v auth_password="'${PG_AUTH_PASSWORD}'" -v notifications_password="'${PG_NOTIFICATIONS_PASSWORD}'" -v posts_password="'${PG_POSTS_PASSWORD}'" -v subscriptions_password="'${PG_SUBSCRIPTIONS_PASSWORD}'" -v author_subscriptions_password="'${PG_AUTHOR_SUBSCRIPTIONS_PASSWORD}'" -v users_password="'${PG_USERS_PASSWORD}'" -f ../migrations/users_up.sql
