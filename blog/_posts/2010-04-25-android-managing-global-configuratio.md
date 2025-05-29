---
layout: post
title: Android - Managing Global Configuration
date: 2010-04-25 00:39:13.000000000 +01:00
type: post
parent_id: "0"
published: true
password: ""
status: publish
categories:
  - Android Development
tags:
  - Configuration Management
  - Mobile software
  - Preferences
  - Settings
  - Smartphones
  - XML
meta:
  _publicize_pending: "1"
  _edit_last: "48492462"
  oc_metadata:
    "{\t\tversion:'1.1',\t\ttags: {'android':
    {\"text\":\"Android\",\"slug\":\"android\",\"source\":{\"_className\":\"SocialTag\",\"url\":\"http://d.opencalais.com/dochash-1/423e44ed-aa17-3c20-a657-f4c9bfe2fb57/SocialTag/1\",\"subjectURL\":null,\"type\":{\"_className\":\"ArtifactType\",\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\"},\"name\":\"Android\",\"makeMeATag\":true,\"importance\":1,\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'mobile-software': {\"text\":\"Mobile
    software\",\"slug\":\"mobile-software\",\"source\":{\"_className\":\"SocialTag\",\"url\":\"http://d.opencalais.com/dochash-1/423e44ed-aa17-3c20-a657-f4c9bfe2fb57/SocialTag/2\",\"subjectURL\":null,\"type\":{\"_className\":\"ArtifactType\",\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\"},\"name\":\"Mobile
    software\",\"makeMeATag\":true,\"importance\":1,\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'smartphones':
    {\"text\":\"Smartphones\",\"slug\":\"smartphones\",\"source\":{\"_className\":\"SocialTag\",\"url\":\"http://d.opencalais.com/dochash-1/423e44ed-aa17-3c20-a657-f4c9bfe2fb57/SocialTag/5\",\"subjectURL\":null,\"type\":{\"_className\":\"ArtifactType\",\"url\":\"http://s.opencalais.com/1/type/tag/SocialTag\",\"name\":\"SocialTag\"},\"name\":\"Smartphones\",\"makeMeATag\":true,\"importance\":1,\"normalizedRelevance\":1},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'xml':
    {\"text\":\"XML\",\"slug\":\"xml\",\"source\":{\"_className\":\"Entity\",\"url\":\"http://d.opencalais.com/genericHasher-1/0ace00c6-2b9f-32c2-8949-82a0f6c6b444\",\"subjectURL\":null,\"type\":{\"_className\":\"ArtifactType\",\"url\":\"http://s.opencalais.com/1/type/em/e/Technology\",\"name\":\"Technology\"},\"name\":\"XML\",\"rawRelevance\":0.192,\"normalizedRelevance\":0.192},\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'configuration-management': {\"text\":\"Configuration
    Management\",\"slug\":\"configuration-management\",\"source\":null,\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'preferences':
    {\"text\":\"Preferences\",\"slug\":\"preferences\",\"source\":null,\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"},
    'settings':
    {\"text\":\"Settings\",\"slug\":\"settings\",\"source\":null,\"bucketName\":\"current\",\"bucketPlacement\":\"auto\",\"_className\":\"Tag\"}}\t}"
  oc_commit_id: http://drone-ah.com/2010/04/25/android-managing-global-configuratio/1272155960
  restapi_import_id: 591d994f7aad5
  original_post_id: "366"
  _wp_old_slug: "366"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:40"
permalink: "/2010/04/25/android-managing-global-configuratio/"
excerpt:
  Accessing configuration / preferences / setttings globally within an
  application is key. The Android API provides an easy way to manage
  configuration but making it globally accessible has a few options. One of
  these options is covered here
---

**The Problem**

Accessing preferences / configuration / settings from Android is actually pretty
straightforward as long as you are in an
[Activity](http://developer.android.com/reference/android/app/Activity.html "Activity"){target="\_blank"}.
To read:

    // PREFS_FILENAME = "nameOfPrefsFile";

    SharedPreferences pref = getSharedPreferences(PREFS_FILENAME,
                                  Context.MODE_PRIVATE);

    String string = pref.getString("key", "default");
    // 1 is the default if key isn't set
    int intValue = pref.getInt("intKey", 1);

    // and so on

[SharedPreference](http://developer.android.com/reference/android/content/SharedPreferences.html "SharedPreferences"){target="\_blank"}s
is the key class. To write, you also need the
[SharedPreferences.Editor](http://developer.android.com/reference/android/content/SharedPreferences.Editor.html){target="\_blank"}
class, as follows:

    // PREFS_FILENAME = "nameOfPrefsFile";
    SharedPreferences pref = getSharedPreferences(PREFS_FILENAME,
                                  Context.MODE_PRIVATE);
    Editor editor = pref.edit();
    editor.putString("key", "value");
    editor.putInt("intKey", 5);

    // Until you call commit, the changes will not
    // be written, so don't forget this step
    editor.commit();

In general however, you will need access to settings in more than one activity
and it seems a bit wasteful to get these bits littered through the application.
Since I am lazy and like to write things just once, I  separated all the prefs
stuff into one class called Settings.

It has a constructor which takes a
[Context](http://developer.android.com/reference/android/content/Context.html "Context"){target="\_blank"}
(We need this to access the SharedPreferences Object). It also has setters and
getters for each property being saved. This example, just saves/retrieves a
username and password.

    import uk.co.kraya.HelloWS;
    import android.content.Context;
    import android.content.SharedPreferences;
    import android.content.SharedPreferences.Editor;

    /**
     * @author Shriram Shri Shrikumar
     *
     * This class stores and manages all the preferences
     * for the application.
     *
     */
    public class Settings {

        private static final String USERNAME_KEY = "username";
        private static final String PASSWORD_KEY = "password";

        private static final String USERNAME_DEFAULT = "username";
        private static final String PASSWORD_DEFAULT = "password";

        private final SharedPreferences settings;

        /**
         * @param act The context from which to pick SharedPreferences
         */
        public Settings (Context act) {
             settings = act.getSharedPreferences(HelloWS.PREFS_NAME, Context.MODE_PRIVATE);
        }

        /**
         * Set the username in the preferences.
         *
         * @param username the username to save into prefs
         */
        public void setUsername(String username) {
            Editor editor = settings.edit();
            editor.putString(USERNAME_KEY, username);
            editor.commit();
        }

        /**
         * @return the username from the prefs
         */
        public String getUsername() {
            return settings.getString(USERNAME_KEY, USERNAME_DEFAULT);
        }

        /**
         *
         * Set the password in the preferences.
         *
         * @param password password to save into prefs
         */
        public void setPassword(String password) {
            Editor editor = settings.edit();
            editor.putString(PASSWORD_KEY, password);
            editor.commit();
        }

        /**
         * @return the password stored in prefs
         */
        public String getPassword() {
            return settings.getString(PASSWORD_KEY, PASSWORD_DEFAULT);
        }

            // Check if there are any stored settings.
            // can be used to automatically load the settings page
            // where necessary
        public boolean hasSettings() {
            // We just check if a username has been set
            return (!settings.getString(USERNAME_KEY, "").equals(""));
        }

    }

Nothing particularly exciting. Now, how do we access this. The Android framework
has a neat little feature that is not very well documented and it involved the
use of the
[Application](http://developer.android.com/reference/android/app/Application.html){target="\_blank"}
class. If you inherit from this class, and point to it in the manifest file, it
will get initialised first before any other objects. This is an ideal place for
bits that need global access. You could use Singletons or static fields but this
works with the framework.

There are two parts to making this work

The application class:

    public class MyApp extends Application {

        private Settings settings;

        @Override
        public void onCreate() {
            settings = new Settings(this);

        }

        public Settings getSettings() {
            return settings;
        }

    }

The onCreate method on MyApp will be called before onCreate on any of the
Activities. The Settings class described above, needs a Context to be passed in.
Lucky for us ;-) Application is also a Context.

You also need to wire it into the AndroidManifest.xml. You need to add the
[android:name](http://developer.android.com/guide/topics/manifest/application-element.html#nm){target="\_blank"}
element into the
[application tag](http://developer.android.com/guide/topics/manifest/application-element.html){target="\_blank"}.

    <application android:name="com.package.MyApp" android:icon="@drawable/icon" android:label="@string/app_name">

Now that is all wired in, accessing the settings object from any activity is
simple:

    MyApp app = (MyApp) getApplicationContext();

    Settings settings = app.getSettings();

Easy - right? While you won\'t be able to access the application subclass
outside of a context, the Setting class, with its local context variable can be
passed around with impunity :-D
