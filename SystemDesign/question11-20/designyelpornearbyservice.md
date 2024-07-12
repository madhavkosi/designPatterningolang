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


## 4. Database Schema

**Places Table**: This table will store information about each place.

- **LocationID** (8 bytes): Uniquely identifies a location.
- **Name** (256 bytes): Name of the place.
- **Latitude** (8 bytes): Latitude of the place.
- **Longitude** (8 bytes): Longitude of the place.
- **Description** (512 bytes): Description of the place.
- **Category** (1 byte): Category of the place (e.g., coffee shop, restaurant, theater, etc.).

**Total Size**: 8 + 256 + 8 + 8 + 512 + 1 = 793 bytes

**Reviews Table**: This table will store reviews for each place.

- **LocationID** (8 bytes): Foreign key referencing the Places table.
- **ReviewID** (4 bytes): Uniquely identifies a review.
- **ReviewText** (512 bytes): Text of the review.
- **Rating** (1 byte): Rating of the place (0-10 stars).

**Total Size**: 8 + 4 + 512 + 1 = 525 bytes

**Photos Table**: This table will store photos for each place and review.

- **PhotoID** (4 bytes): Uniquely identifies a photo.
- **LocationID** (8 bytes): Foreign key referencing the Places table.
- **ReviewID** (4 bytes): Foreign key referencing the Reviews table (can be NULL for place photos).
- **PhotoURL** (256 bytes): URL of the photo.

**Total Size**: 4 + 8 + 4 + 256 = 272 bytes


## 5 Api contract

| **Function**          | **Parameters**                                                                                                                                                                       | **Returns**                                                |
|-----------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------|
| **search**            | api_dev_key (string), search_terms (string), user_location (string), radius_filter (number), maximum_results_to_return (number), category_filter (string), sort (number), page_token (string) | JSON with a list of places (name, address, category, rating, thumbnail) |
| **add_place**         | api_dev_key (string), name (string), latitude (number), longitude (number), description (string), category (string)                                                                   | JSON confirming place creation                             |
| **get_place_details** | api_dev_key (string), location_id (number)                                                                                                                                           | JSON with detailed information about the place             |
| **add_review**        | api_dev_key (string), location_id (number), review_text (string), rating (number), photos (array of strings)                                                                          | JSON confirming review submission                          |
| **get_reviews_for_place** | api_dev_key (string), location_id (number)                                                                                                                                           | JSON with a list of reviews for the specified place        |

