## Designing a Video Sharing Service (Similar to YouTube)

### Why YouTube?
YouTube is one of the most popular video-sharing platforms globally. It allows users to upload, view, share, rate, and report videos, as well as add comments. This makes it an ideal model for designing a similar service.

### Requirements and Goals of the System

#### **Functional Requirements:**
1. Users can upload videos.
2. Users can share and view videos.
3. Users can search for videos by title.
4. System records video stats (likes/dislikes, total views).
5. Users can add and view comments on videos.

#### **Non-Functional Requirements:**
1. Highly reliable: No uploaded video should be lost.
2. Highly available: Prioritize availability over consistency; temporary video unavailability is acceptable.
3. Real-time experience: Ensure smooth, lag-free video playback.

**Out of Scope:**
- Excludes features like recommendations, popular videos, channels, subscriptions, watch later, and favorites.


### Capacity Estimation and Constraints

**User and Activity Estimates:**
- Total users: 1.5 billion
- Daily active users: 800 million
- Average video views per user per day: 5
- Total video views per second: 46,000 (800M * 5 / 86,400)

**Upload Activity:**
- Upload:view ratio: 1:200
- Video uploads per second: 230 (46,000 / 200)

**Storage Estimates:**
- Video upload rate: 500 hours of video per minute
- Storage needed per minute of video: 50MB
- Total storage needed per minute: 1,500 GB (500 * 60 * 50)
- Storage needed per second: 25 GB (1,500 / 60)

**Bandwidth Estimates:**
- Bandwidth for uploads per minute: 300GB (500 * 60 * 10MB) (assumung 10MB of upload)
- Bandwidth for uploads per second: 5GB (300 / 60)
- Bandwidth for views (upload:view ratio of 1:200): 1TB/s

These estimates ignore video compression and replication, which would adjust real numbers.
- Upload:view ratio: 1:200
- Video uploads per second: 2