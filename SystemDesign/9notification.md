## Designing a Notification System

### Overview

**Purpose of Notifications:**
- Alerts users with important information such as breaking news, product updates, events, and offerings.
- Indispensable part of daily life.

### Types of Notifications

1. **Mobile Push Notification:**
   - Delivered directly to the user's mobile device.
   - Often used for immediate or urgent information.

2. **SMS Message:**
   - Sent as text messages to the user's phone.
   - Suitable for concise and urgent messages.

3. **Email:**
   - Sent to the user's email address.
   - Ideal for detailed information, newsletters, and non-urgent updates.

## Step 1 - Understand the Problem and Scope

**Key Points:**
- **Scalability:** Handle millions of daily notifications.
- **Notification Types:** Push notification, SMS, email.
- **System Nature:** Soft real-time, slight delays acceptable.
- **Supported Devices:** iOS, Android, laptops/desktops.
- **Triggers:** Client apps, server-side scheduling.
- **Opt-Out:** Users can opt-out.
- **Daily Volume:** 10M push, 1M SMS, 5M emails.

## Step 2 - Propose High-Level Design and Get Buy-In

**Overview:**
The goal is to create a high-level design supporting various notification types including iOS push notifications, Android push notifications, SMS messages, and Emails. This design will be structured as follows:
1. Different types of notifications
2. Contact info gathering flow
3. Notification sending/receiving flow

### Different Types of Notifications

1. **iOS Push Notification:**
   - **Components:**
     - **Provider:** Builds and sends notification requests to Apple Push Notification Service (APNS).
       - **Device Token:** A unique identifier for sending push notifications to a specific device.
       - **Payload:** A JSON dictionary containing the notification’s content.
         - **Example Payload:**
           ```json
           {
             "aps": {
               "alert": {
                 "title": "Game Request",
                 "body": "Bob wants to play chess",
                 "action-loc-key": "PLAY"
               },
               "badge": 5
             }
           }
           ```
     - **APNS:** Apple’s remote service for propagating push notifications to iOS devices.
     - **iOS Device:** The end client that receives and displays the push notifications.

2. **Android Push Notification:**
   - **Components:**
     - **Provider:** Uses Firebase Cloud Messaging (FCM) to send push notifications.
     - **Android Device:** The end client that receives and displays the push notifications.

3. **SMS Message:**
   - **Components:**
     - Utilizes third-party SMS services such as Twilio or Nexmo for sending messages.
     - These services handle the delivery of SMS to users' mobile phones.

4. **Email:**
   - **Components:**
     - Uses commercial email services like SendGrid or Mailchimp.
     - These services offer better delivery rates and data analytics compared to setting up own email servers.

<p float="left">
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/email.webp" width="300" />
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/android.webp" width="300" />
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/ios.webp" width="300" />
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/sms.webp" width="300" />
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/combined.webp" width="150" />
</p>

### Contact Info Gathering Flow

**Process:**
- Collect user contact information (device tokens, phone numbers, email addresses) when a user installs the app or signs up.
- API servers handle the collection and storage of this information in a database.
  
**Database Structure:**
- **User Table:** Stores user information including email addresses and phone numbers.
- **Device Table:** Stores device tokens. A user can have multiple devices, allowing for notifications to be sent to all user devices.

<p float="left">
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/notifdb.svg" width="500" />
</p>

### Notification Sending/Receiving Flow

**Initial Design:**

**Components:**
- **Service 1 to N:** Different services (e.g., billing, shopping) that trigger notification sending events.
- **Notification System:** Centralized server that provides APIs for services, constructs notification payloads, and integrates with third-party services.
- **Third-Party Services:** Responsible for delivering notifications to user devices.

**Challenges:**
- **Single Point of Failure (SPOF):** A single notification server is a potential SPOF.
- **Scalability:** Handling all notifications on one server makes scaling difficult.
- **Performance Bottleneck:** Processing and sending notifications can be resource-intensive, leading to potential overload during peak times.

<p float="left">
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/firstversion.webp" width="500" />
</p>

**Improved High-Level Design:**

**Improvements:**
- Move database and cache outside of the notification server.
- Add more notification servers and enable automatic horizontal scaling.
- Introduce message queues to decouple system components.

**Components:**
- **Service 1 to N:** Represent different services that send notifications via provided APIs.
- **Notification Servers:**
  - Provide APIs for services to send notifications.
  - Conduct basic validations (e.g., email, phone number verification).
  - Fetch data from cache or database for rendering notifications.
  - Send notification data to message queues for parallel processing.
  - **Example API for Sending Email:**
    ```http
    POST https://api.example.com/v1/email/send
    {
      "to": [{"user_id": 123456}],
      "from": {"email": "from_address@example.com"},
      "subject": "Hello World!",
      "content": [{"type": "text/plain", "value": "Hello, World!"}]
    }
    ```
- **Cache:** Stores frequently accessed data like user info, device info, and notification templates.
- **Database:** Stores user data, notification settings, etc.
- **Message Queues:** Buffer high volumes of notifications, with distinct queues for each notification type to ensure independence.
- **Workers:** Pull notification events from message queues and send them to third-party services.
- **Third-Party Services:** Deliver notifications to user devices.
- **User Devices:** Receive and display the notifications.

<p float="left">
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/enhancedversion.webp" width="500" />
</p>

**Notification Flow:**
1. A service calls APIs provided by notification servers to send notifications.
2. Notification servers fetch metadata (user info, device token, notification settings) from cache or database.
3. Notification event is sent to the appropriate message queue.
4. Workers pull events from the queues.
5. Workers send notifications to third-party services.
6. Third-party services deliver notifications to user devices.

## Step 3 - Design Deep Dive

**Overview:**
In this step, we will explore the following aspects in-depth:
1. Reliability
2. Additional components and considerations
3. Updated design

### Reliability

1. **Preventing Data Loss:**
   - Notifications must not be lost; they can be delayed or re-ordered.
   - Persist notification data in a database.
   - Implement a retry mechanism.
   - Include a notification log database for data persistence.

2. **Ensuring Recipients Receive a Notification Exactly Once:**
   - Exact once delivery is not guaranteed due to the distributed nature.
   - Implement a deduplication (dedupe) mechanism:
     - Check if the event ID has been seen before; if yes, discard the notification, otherwise send it.
   - Reference material [5] for more details on delivery mechanisms.

### Additional Components and Considerations

1. **Notification Template:**
   - Templates avoid building each notification from scratch.
   - Benefits include consistency, reduced error margin, and time-saving.
   - Example template:
     ```
     BODY: You dreamed of it. We dared it. [ITEM NAME] is back — only until [DATE].
     CTA: Order Now. Or, Save My [ITEM NAME]
     ```

2. **Notification Settings:**
   - Users can control their notification preferences.
   - Stored in a notification setting table with fields:
     - `user_id` (bigInt)
     - `channel` (varchar: push notification, email, or SMS)
     - `opt_in` (boolean: opt-in to receive notifications)

3. **Rate Limiting:**
   - Limit the number of notifications sent to avoid overwhelming users.
   - Prevent users from turning off notifications due to frequent messages.

4. **Retry Mechanism:**
   - Retry sending notifications if a third-party service fails.
   - If problems persist, alert developers.

5. **Security in Push Notifications:**
   - Use `appKey` and `appSecret` for securing push notification APIs.
   - Authenticate clients before allowing them to send push notifications.

6. **Monitor Queued Notifications:**
   - Monitor the total number of queued notifications.
   - Increase workers if the number of queued notifications is high to avoid

 delays.

7. **Events Tracking:**
   - Track metrics like open rate, click rate, and engagement.
   - Integrate with analytics services for better understanding of customer behavior.

### Updated Design

**Enhanced Features:**
- **Notification Servers:**
  - Equipped with authentication and rate-limiting features.
  - Implement retry mechanisms for failed notifications.
  - Use notification templates for efficient and consistent creation.
- **Monitoring and Tracking Systems:**
  - Added for system health checks and future improvements.
  
**Flow Overview:**
1. **Reliability Measures:**
   - Notification log database for data persistence.
   - Retry mechanisms to ensure delivery.

2. **Notification Creation and Sending:**
   - Use templates for creating notifications.
   - Respect user notification settings and apply rate limits.
   - Implement deduplication logic to minimize duplicates.

3. **Security and Monitoring:**
   - Secure APIs with `appKey` and `appSecret`.
   - Monitor queued notifications and add more workers if needed.
   - Track user engagement metrics and integrate with analytics services.

The updated design includes improvements in reliability, user control, security, monitoring, and tracking, ensuring a robust and efficient notification system.

<p float="left">
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/enhancedversion.webp" width="800" />
</p>