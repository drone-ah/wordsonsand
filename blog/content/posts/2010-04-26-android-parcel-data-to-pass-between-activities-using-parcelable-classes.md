---
categories:
  - Android Development
date: "2010-04-26T21:46:31Z"
meta:
  _edit_last: "48492462"
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:54:40"
  _publicize_pending: "1"
  _wp_old_slug: "378"
  original_post_id: "378"
  restapi_import_id: 591d994f7aad5
parent_id: "0"
password: ""
status: publish
tags:
  - android-development
  - activity
  - intent
  - java
  - parcel
  - parcelable
title: Android - Parcel data to pass between Activities using Parcelable classes
url: /2010/04/26/android-parcel-data-to-pass-between-activities-using-parcelable-classes/
---

Passing data between activities on android is unfortunately, not as simple as
passing in parameters. What we need to to do is tag these onto the intent. If
the information we need to pass across is a simple object like a String or
Integer, this is easy enough.

```java
String strinParam = "String Parameter";
Integer intParam = 5;

Intent i = new Intent(this, MyActivity.class);
i.putExtra("uk.co.kraya.stringParam", stringParam);
i.putExtra("uk.co.kraya.intParam", intParam);

startActivity(i);
```

<!--more-->

Passing in custom objects is a little more complicated. You could just mark the
class
as [Serializable](http://java.sun.com/javase/6/docs/api/java/io/Serializable.html)
and let Java take care of this. However, on the android, there is a serious
performance hit that comes with using Serializable. The solution is to
use [Parcelable](http://developer.android.com/reference/android/os/Parcelable.html).

```java
package uk.co.kraya.android.demos.Parcelable;

import android.os.Parcel;
import android.os.Parcelable;

/**
 * @author Shriram Shri Shrikumar
 *
 * A basic object that can be parcelled to
 * transfer between objects
 *
 */
public class ObjectA implements Parcelable {

    private String strValue;
    private Integer intValue;

    /**
     * Standard basic constructor for non-parcel
     * object creation
     */
    public ObjectA() { ; };

    /**
     *
     * Constructor to use when re-constructing object
     * from a parcel
     *
     * @param in a parcel from which to read this object
     */
    public ObjectA(Parcel in) {
        readFromParcel(in);
    }

    /**
     * standard getter
     *
     * @return strValue
     */
    public String getStrValue() {
        return strValue;
    }

    /**
     * Standard setter
     *
     * @param strValue
     */
    public void setStrValue(String strValue) {
        this.strValue = strValue;
    }

    /**
     * standard getter
     *
     * @return
     */
    public Integer getIntValue() {
        return intValue;
    }

    /**
     * Standard setter
     *
     * @param intValue
     */
    public void setIntValue(Integer intValue) {
        this.intValue = intValue;
    }

    @Override
    public int describeContents() {
        return 0;
    }

    @Override
    public void writeToParcel(Parcel dest, int flags) {

        // We just need to write each field into the
        // parcel. When we read from parcel, they
        // will come back in the same order
        dest.writeString(strValue);
        dest.writeInt(intValue);
    }

    /**
     *
     * Called from the constructor to create this
     * object from a parcel.
     *
     * @param in parcel from which to re-create object
     */
    private void readFromParcel(Parcel in) {

        // We just need to read back each
        // field in the order that it was
        // written to the parcel
        strValue = in.readString();
        intValue = in.readInt();
    }

    /**
     *
     * This field is needed for Android to be able to
     * create new objects, individually or as arrays.
     *
     * This also means that you can use use the default
     * constructor to create the object and use another
     * method to hyrdate it as necessary.
     *
     * I just find it easier to use the constructor.
     * It makes sense for the way my brain thinks ;-)
     *
     */
    public static final Parcelable.Creator CREATOR =
        new Parcelable.Creator() {
            public ObjectA createFromParcel(Parcel in) {
                return new ObjectA(in);
            }

            public ObjectA[] newArray(int size) {
                return new ObjectA[size];
            }
        };

}
```

The intricacies of the class is described in the code above. There is now one
more special case. What if you have an object that references another object.
Clearly, they would both need to be Parcelable, but how would be integrate them.
ObjectB shows a parcelable embedded in another parcelable...

```java
package uk.co.kraya.android.demos.Parcelable;

import android.os.Parcel;
import android.os.Parcelable;

public class ObjectB implements Parcelable {

    private ObjectA obj;
    private Long longVal;

    public ObjectB() { ; }

    public ObjectA getObj() {
        return obj;
    }

    /**
     *
     * Constructor to use when re-constructing object
     * from a parcel
     *
     * @param in a parcel from which to read this object
     */
    public ObjectB(Parcel in) {
        readFromParcel(in);
    }

    public void setObj(ObjectA obj) {
        this.obj = obj;
    }

    public Long getLongVal() {
        return longVal;
    }

    public void setLongVal(Long longVal) {
        this.longVal = longVal;
    }

    @Override
    public int describeContents() {
        return 0;
    }

    @Override
    public void writeToParcel(Parcel dest, int flags) {

        // The writeParcel method needs the flag
        // as well - but thats easy.
        dest.writeParcelable(obj, flags);

        // Same as in ObjectA
        dest.writeLong(longVal);
    }

    /**
     *
     * Called from the constructor to create this
     * object from a parcel.
     *
     * @param in parcel from which to re-create object
     */
    private void readFromParcel(Parcel in) {

        // readParcelable needs the ClassLoader
        // but that can be picked up from the class
        // This will solve the BadParcelableException
        // because of ClassNotFoundException
        obj = in.readParcelable(ObjectA.class.getClassLoader());

        // The rest is the same as in ObjectA
        longVal = in.readLong();
    }

    /**
     *
     * This field is needed for Android to be able to
     * create new objects, individually or as arrays.
     *
     * This also means that you can use use the default
     * constructor to create the object and use another
     * method to hyrdate it as necessary.
     *
     * I just find it easier to use the constructor.
     * It makes sense for the way my brain thinks ;-)
     *
     */
    public static final Parcelable.Creator CREATOR =
        new Parcelable.Creator() {
            public ObjectB createFromParcel(Parcel in) {
                return new ObjectB(in);
            }

            public ObjectB[] newArray(int size) {
                return new ObjectB[size];
            }
        };
}
```

When writing the parcel, we need to pass in the flags - which is easy enough.
When reading the parcel, we need the classloader, which can be picked up from
destination class of the parcelable. Again easy!

Finally, passing a parcelable object to an intent

```java
ObjectA obj = new ObjectA();

// Set values etc.

Intent i = new Intent(this, MyActivity.class);
i.putExtra("com.package.ObjectA", obj);

startActivity(i);
```

Almost too easy - right?

and to read the values,

```java
public class MyActivity extends Activity {

    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

        Bundle b = getIntent().getExtras();
        ObjectA obj =
            b.getParcelable("com.package.ObjectA");
    }

}
```

It it was any easier - we'd all be out of a job ;-)
