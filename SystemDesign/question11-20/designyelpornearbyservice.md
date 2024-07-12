# Designing a Yelp-like Service or Nearby Friends

## 1. Introduction

A Yelp-like service allows users to search for nearby places such as restaurants, theaters, shopping malls, etc., and to add/view reviews of these places. This service can be considered a proximity server used to discover nearby attractions.

## 2. Requirements and Goals

### Functional Requirements

1. Users can add, delete, and update information about places.
2. Users can find nearby places given their location (longitude/latitude) within a specified radius.
3. Users can add reviews/feedback about a place, including pictures, text, and a rating.

### Non-functional Requirements

1. Real-time search experience with minimal latency.
2. Support for a heavy search load, with a higher volume of search requests compared to place updates.

## 3. Scale Estimation

1. **Places**: 500M places
2. **Queries per Second (QPS)**: 100K
3. **Annual Growth**: 20% increase in places and QPS
