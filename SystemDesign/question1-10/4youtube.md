
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


### System APIs - 

**API Definitions**

**uploadVideo**
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

**searchVideo**
- **Endpoint**: `GET /api/searchVideo`
- **Parameters**:
  - `api_dev_key (string)`: API developer key.
  - `search_query (string)`: Search terms.
  - `user_location (string)`: Optional user location.
  - `maximum_videos_to_return (number)`: Max results per request.
  - `page_token (string)`: Page token for results.
- **Returns**:
  - `JSON` with video list including title, thumbnail, creation date, and view count.

**streamVideo**
- **Endpoint**: `GET /api/streamVideo`
- **Parameters**:
  - `api_dev_key (string)`: API developer key.
  - `video_id (string)`: Video identifier.
  - `offset (number)`: Time in seconds from start.
  - `codec (string)`: Video codec info.
  - `resolution (string)`: Video resolution.
- **Returns**:
  - `STREAM`: Video chunk from given offset.

### Database Design
To design a scalable database schema for a video-sharing platform, we'll need to consider partitioning, indexing, and optimizing data access patterns to ensure the system can handle large volumes of data and high traffic. Here’s an updated schema that includes these considerations:

**Scalable Database Schema Design**

**Tables**:

1. **Users**
   - `user_id`: `BIGINT` (Primary Key, AUTO_INCREMENT)
   - `username`: `VARCHAR(255)`
   - `email`: `VARCHAR(255)`
   - `password_hash`: `VARCHAR(255)`
   - `created_at`: `TIMESTAMP`
   - `updated_at`: `TIMESTAMP`

2. **Videos**
   - `video_id`: `BIGINT` (Primary Key, AUTO_INCREMENT)
   - `user_id`: `BIGINT` (Foreign Key references Users(user_id))
   - `title`: `VARCHAR(255)`
   - `description`: `TEXT`
   - `tags`: `VARCHAR(255)`
   - `category_id`: `BIGINT`
   - `default_language`: `VARCHAR(50)`
   - `recording_details`: `TEXT`
   - `file_path`: `VARCHAR(255)`
   - `thumbnail_path`: `VARCHAR(255)`
   - `created_at`: `TIMESTAMP`
   - `updated_at`: `TIMESTAMP`
   - `views_count`: `BIGINT`
   - `likes_count`: `BIGINT`
   - `dislikes_count`: `BIGINT`
   - `comments_count`: `BIGINT`

3. **Categories**
   - `category_id`: `BIGINT` (Primary Key, AUTO_INCREMENT)
   - `name`: `VARCHAR(255)`
   - `description`: `TEXT`
   - `created_at`: `TIMESTAMP`
   - `updated_at`: `TIMESTAMP`

4. **Comments**
   - `comment_id`: `BIGINT` (Primary Key, AUTO_INCREMENT)
   - `video_id`: `BIGINT` (Foreign Key references Videos(video_id))
   - `user_id`: `BIGINT` (Foreign Key references Users(user_id))
   - `comment_text`: `TEXT`
   - `likes_count`: `BIGINT`
   - `dislikes_count`: `BIGINT`
   - `created_at`: `TIMESTAMP`
   - `updated_at`: `TIMESTAMP`

5. **Likes**
   - `like_id`: `BIGINT` (Primary Key, AUTO_INCREMENT)
   - `user_id`: `BIGINT` (Foreign Key references Users(user_id))
   - `video_id`: `BIGINT` (Foreign Key references Videos(video_id), nullable if comment_id is present)
   - `comment_id`: `BIGINT` (Foreign Key references Comments(comment_id), nullable if video_id is present)
   - `is_like`: `BOOLEAN`
   - `created_at`: `TIMESTAMP`
   - `updated_at`: `TIMESTAMP`

6. **VideoViews**
   - `view_id`: `BIGINT` (Primary Key, AUTO_INCREMENT)
   - `video_id`: `BIGINT` (Foreign Key references Videos(video_id))
   - `user_id`: `BIGINT` (Foreign Key references Users(user_id))
   - `viewed_at`: `TIMESTAMP`


### Very  High-Level Design
**System Components Overview**
- **Components** (Refer to Figure 3):
  - **Client**: Access YouTube on computer, mobile phone, or smartTV.
  - **CDN**: Stores and streams videos.
  - **API Servers**: Handle all requests except video streaming (e.g., feed recommendation, video upload URL generation, metadata updates, user signup).

<p float="left">
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/youtube1.webp" width="500" />
</p>

**Key Flows of Interest**

- **Video Uploading Flow**
- **Video Streaming Flow**

**Video Uploading Flow**

**Components**

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
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/youtube2.webp" width="500" />
</p>




**Video Uploading Flow Breakdown**

**Flow A: Upload the Actual Video (Refer to Figure 5)**

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

**Flow B: Update the Metadata (Refer to Figure 6)**

- **Parallel Process**: While uploading, client sends metadata update request to API servers.
- **Update**: API servers update metadata cache and database.

**Video Streaming Flow**

- **Concept**: Streaming allows immediate and continuous playback without waiting for the entire video to download.
- **Streaming Protocols**:
  - **MPEG–DASH**: Dynamic Adaptive Streaming over HTTP.
  
  - **Apple HLS**: HTTP Live Streaming.
  - **Microsoft Smooth Streaming**
  - **Adobe HTTP Dynamic Streaming (HDS)**
  - **Importance**: Different protocols support various encodings and playback players.

  **Dynamic Adaptive Streaming over HTTP (DASH)**
  - **Developed by:** MPEG (Moving Picture Experts Group)
  - **Use Case:** Adaptive streaming for various media types.
  - **Key Features:**
    - HTTP-based
    - Supports adaptive bitrate streaming
    - Device and vendor-agnostic
    - Media content is divided into segments

- **CDN Streaming**: Videos streamed from the nearest edge server to minimize latency (Refer to Figure 7).

 **Summary**

- **Cloud Services**: Use CDN and blob storage for scalability and cost-effectiveness.
- **System Design**: Focus on high-level components and their interactions.
- **Key Flows**:
  - **Video Uploading**: Involves original storage, transcoding, metadata updates, and CDN distribution.
  - **Video Streaming**: Utilizes appropriate streaming protocols and CDN for efficient delivery.
- **Next Steps**: Explore detailed designs for video uploading and streaming flows based on high-level architecture.

---
### Design deep dive

**Video Transcoding**
**Importance**
- **Storage Efficiency**: Raw videos consume substantial storage; transcoding reduces the file size significantly.
- **Device Compatibility**: Ensures videos play smoothly across different devices and browsers by converting them into universally supported formats.
- **Adaptive Streaming**: Adjusts video quality dynamically based on user’s network conditions, providing a smooth viewing experience regardless of bandwidth.


**Directed Acyclic Graph (DAG) Model for Video Transcoding**

**Purpose**
- **Efficiency**: Handles computationally expensive and time-consuming video transcoding tasks.
- **Flexibility**: Supports varied video processing requirements from different content creators.
- **Parallelism**: Enables high parallelism in processing tasks to optimize performance.

**Key Concepts**
- **Task Abstraction**: Allows client programmers to define tasks for specific video processing needs.
- **Sequential and Parallel Execution**: Tasks are defined in stages, enabling them to be executed either sequentially or in parallel.

**DAG Components**
- **Video Splitting**: The original video is divided into three primary components:
  - **Video**: The visual content.
  - **Audio**: The sound component.
  - **Metadata**: Additional data providing context and information about the video.

<p float="left">
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/youtube1.svg" width="500" />
    <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/youtube2.svg" width="300" />
</p>

**Tasks in the DAG Model**
1. **Inspection**:
   - Ensures video quality is high and the video files are not malformed.
   
2. **Video Encodings**:
   - Converts videos to support different resolutions, codecs, and bitrates.
   - Examples include encoding files to various formats as shown in Figure 9.
   
3. **Thumbnail Generation**:
   - Thumbnails can be either uploaded by users or automatically generated by the system.
   
4. **Watermarking**:
   - Adds an image overlay on top of the video to include identifying information.

By adopting a DAG model, video transcoding systems achieve the necessary flexibility and efficiency to handle a wide range of video processing tasks, ensuring high performance and quality in video delivery.


**Video Transcoding Architecture**

**Overview**
The architecture leverages cloud services and consists of six main components: preprocessor, DAG scheduler, resource manager, task workers, temporary storage, and encoded video output.

<p float="left">
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/youtube3.svg" width="500" />
</p>

**Components**

**Preprocessor**

1. **Video Splitting**: Splits video into smaller GOP units.
2. **Legacy Support**: Ensures compatibility with older devices/browsers.
3. **DAG Generation**: Creates DAG from client configurations.
4. **Cache Data**: Caches segmented videos and metadata for reliability and retries.

**DAG Scheduler**
   - **Function**: Splits a DAG into stages of tasks.
   - **Process**:
     - Places tasks in the task queue of the resource manager.
     - Example: Original video split into three stages (video, audio, metadata), further divided into tasks like video encoding and thumbnail generation (Figure 15).

**Resource Manager**
   - **Responsibilities**: Manages resource allocation efficiency.
   - **Components**:
     - **Task Queue**: Contains tasks to be executed.
     - **Worker Queue**: Contains worker utilization info.
     - **Running Queue**: Tracks currently running tasks and workers.
     - **Task Scheduler**: Picks optimal task/worker, instructs task worker, and manages task execution flow.
   - **Process** (Figure 17):
     - Selects the highest priority task from the task queue.
     - Selects the optimal task worker from the worker queue.
     - Instructs the task worker to execute the job.
     - Updates the running queue with task/worker info.
     - Removes the job from the running queue once completed.

<p float="left">
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/youtube5.svg" width="500" />
</p>

4. **Task Workers**
   - **Function**: Execute tasks defined in the DAG.
   - **Variety**: Different workers may run different tasks (Figure 19).

5. **Temporary Storage**
   - **Usage**: Stores data based on type, size, access frequency, and lifespan.
   - **Metadata**: Cached in memory due to small size and frequent access.
   - **Video/Audio Data**: Stored in blob storage.
   - **Data Lifecycle**: Data in temporary storage is freed up post-processing.

6. **Encoded Video**
   - **Output**: The final result of the encoding pipeline.
   - **Example**: Encoded file such as funny_720p.mp4 (Figure 21).



### System Optimizations

#### Speed Optimization

1. **Parallelize Video Uploading**:
   - **Chunk Uploads**: Split videos into smaller chunks by GOP alignment for faster, resumable uploads.
   - **Implementation**: Clients handle splitting to improve upload speed (Figure 22, Figure 23).

2. **Global Upload Centers**:
   - **CDN as Upload Centers**: Set up multiple upload centers globally to reduce upload time.
   - **Example**: North America and Asia upload centers for users in respective regions (Figure 24).

3. **Parallelism Everywhere**:
   - **Loosely Coupled System**: Use message queues to decouple system components.
   - **Benefit**: Enables encoding to proceed without waiting for download completion, enhancing parallelism (Figure 25, Figure 26).

#### Safety Optimization

1. **Pre-signed Upload URLs**:
   - **Purpose**: Ensure only authorized users upload videos to correct locations.
   - **Flow**:
     1. Client requests a pre-signed URL from API servers.
     2. API servers respond with the URL.
     3. Client uses the URL to upload the video (Figure 27).

<p float="left">
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/youtube7.webp" width="300" />
    <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/youtube8.svg" width="300" />

</p>


### Video Deduplication
- **Purpose**: Reduce storage space, improve caching, decrease network usage, and save energy.
- **Method**: Inline deduplication during upload using video matching algorithms (e.g., Block Matching, Phase Correlation).
- **Process**:
  - Stop upload if duplicate found and use existing copy.
  - Continue upload if new video is higher quality.
  - Upload only missing parts if video is a subpart of an existing video.