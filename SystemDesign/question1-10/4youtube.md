
### Why YouTube?
YouTube is one of the most popular video-sharing platforms globally. It allows users to upload, view, share, rate, and report videos, as well as add comments. This makes it an ideal model for designing a similar service.

### Requirements and Goals of the System

**Functional Requirements:**
1. Users can upload videos.
2. Users can share and view videos.
3. Users can search for videos by title.
4. System records video stats (likes/dislikes, total views).
5. Users can add and view comments on videos.

**Non-Functional Requirements:**
1. Highly reliable: No uploaded video should be lost.
2. Highly available: Prioritize availability over consistency; temporary video unavailability is acceptable.
3. Real-time experience: Ensure smooth, lag-free video playback.

**Out of Scope:**
- Excludes features like recommendations, popular videos, channels, subscriptions, watch later, and favorites.


### Capacity Estimation and Constraints
**Load Estimate**

  **User and Activity Estimates:**
  - Total users: 1.5 billion
  - Daily active users: 800 million
  - Average video views per user per day: 5
  - Total video views per second: 46,000 (800M * 5 / 86,400)

  **Upload Activity:**
  - Upload:view ratio: 1:200
  - Video uploads per second: 230 (46,000 / 200)

  **Likes/Dislikes Estimates:**
  - **Interaction ratio (likes/dislikes)**: Assuming 10% of views
  - **Likes/Dislikes per second**: 4,600 (46,000 * 0.10)

  **comment Estimates:**
  - **Interaction ratio (comment ration)**: Assuming 5% of views
  - **Likes/Dislikes per second**: 23,00 (46,000 * 0.50)

**Storage Estimates:**
  - Assume Average video size is 300 MB
  - Storage Needed (300MB * 230)*86400 =5 PB/day for video
  - For Like 4600/second * 1 Kb *86400 =49GB Per/Day
  - For Comment 2300/second * 2 Kb *86400 =49GB Per/Day

**Bandwidth Estimates:**
 
- Bandwidth for uploads per second: 69GB/s (300 MB *230)
- Bandwidth for views (upload:view ratio of 1:200): 13.8 TB/s (300 MB * 230 * 200)


# System APIs - 

## API Definitions

### uploadVideo
- **Endpoint**: `POST /api/uploadVideo`
- **Parameters**:
  - `api_dev_key (string)`: API developer key.
  - `video_title (string)`: Title of the video.
  - `video_description (string)`: Optional description.
  - `tags (string[])`: Optional tags.
  - `category_id (string)`: Video category.
  - `default_language (string)`: Language of the video.
  - `recording_details (string)`: Recording location.
  - `video_contents (stream)`: Video file.
- **Returns**: 
  - `HTTP 202` on successful upload.
  - Notification email with video link after encoding.
  - Queryable API for upload status.

### searchVideo
- **Endpoint**: `GET /api/searchVideo`
- **Parameters**:
  - `api_dev_key (string)`: API developer key.
  - `search_query (string)`: Search terms.
  - `user_location (string)`: Optional user location.
  - `maximum_videos_to_return (number)`: Max results per request.
  - `page_token (string)`: Page token for results.
- **Returns**:
  - `JSON` with video list including title, thumbnail, creation date, and view count.

### streamVideo
- **Endpoint**: `GET /api/streamVideo`
- **Parameters**:
  - `api_dev_key (string)`: API developer key.
  - `video_id (string)`: Video identifier.
  - `offset (number)`: Time in seconds from start.
  - `codec (string)`: Video codec info.
  - `resolution (string)`: Video resolution.
- **Returns**:
  - `STREAM`: Video chunk from given offset.
  
# Step 2 - Propose High-Level Design and Get Buy-In

## Leveraging Cloud Services

- **Recommendation**: Use existing cloud services like CDN and blob storage instead of building from scratch.
- **Rationale**:
  - **Time Efficiency**: System design interviews focus on choosing the right technology, not on building everything from scratch.
  - **Complexity and Cost**: Building scalable blob storage or CDN is complex and costly. Large companies like Netflix and Facebook use cloud services (e.g., Amazon’s cloud services, Akamai’s CDN).

## System Components Overview

- **Components** (Refer to Figure 3):
  - **Client**: Access YouTube on computer, mobile phone, or smartTV.
  - **CDN**: Stores and streams videos.
  - **API Servers**: Handle all requests except video streaming (e.g., feed recommendation, video upload URL generation, metadata updates, user signup).

## Key Flows of Interest

- **Video Uploading Flow**
- **Video Streaming Flow**

## Video Uploading Flow

### Components (Refer to Figure 4)

1. **User**: Watches YouTube on various devices.
2. **Load Balancer**: Distributes requests among API servers.
3. **API Servers**: Process all user requests except video streaming.
4. **Metadata DB**: Stores video metadata, sharded and replicated for performance and availability.
5. **Metadata Cache**: Caches video metadata and user objects for better performance.
6. **Original Storage**: Blob storage for original videos.
7. **Transcoding Servers**: Convert video formats to support different devices and bandwidths.
8. **Transcoded Storage**: Blob storage for transcoded video files.
9. **CDN**: Caches and streams videos.
10. **Completion Queue**: Stores video transcoding completion events.
11. **Completion Handler**: Workers that update metadata cache and database upon transcoding completion.

<p float="left">
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/youtube1.webp" width="500" />
</p>

### Video Uploading Flow Breakdown

#### Flow A: Upload the Actual Video (Refer to Figure 5)

1. **Upload**: Videos uploaded to original storage.
2. **Transcoding**: Servers fetch and transcode videos.
3. **Post-Transcoding**:
   - **Parallel Steps**:
     - **Transcoded Storage**: Store transcoded videos.
     - **Completion Queue**: Queue transcoding completion events.
   - **Distribution**:
     - **CDN**: Distribute transcoded videos.
     - **Completion Handler**: Update metadata DB and cache from queue events.
4. **Notification**: API servers inform client of successful upload and readiness for streaming.

#### Flow B: Update the Metadata (Refer to Figure 6)

- **Parallel Process**: While uploading, client sends metadata update request to API servers.
- **Update**: API servers update metadata cache and database.

## Video Streaming Flow

- **Concept**: Streaming allows immediate and continuous playback without waiting for the entire video to download.
- **Streaming Protocols**:
  - **MPEG–DASH**: Dynamic Adaptive Streaming over HTTP.
  - **Apple HLS**: HTTP Live Streaming.
  - **Microsoft Smooth Streaming**
  - **Adobe HTTP Dynamic Streaming (HDS)**
  - **Importance**: Different protocols support various encodings and playback players.
- **CDN Streaming**: Videos streamed from the nearest edge server to minimize latency (Refer to Figure 7).

## Summary

- **Cloud Services**: Use CDN and blob storage for scalability and cost-effectiveness.
- **System Design**: Focus on high-level components and their interactions.
- **Key Flows**:
  - **Video Uploading**: Involves original storage, transcoding, metadata updates, and CDN distribution.
  - **Video Streaming**: Utilizes appropriate streaming protocols and CDN for efficient delivery.
- **Next Steps**: Explore detailed designs for video uploading and streaming flows based on high-level architecture.

---

<p float="left">
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/youtube2.webp" width="500" />
</p>


