---
categories:
- Artificial Intelligence
date: "2016-08-09T13:16:11Z"
meta:
  _elasticsearch_data_sharing_indexed_on: "2024-11-18 14:55:01"
  _publicize_job_id: "5186401855"
  _rest_api_client_id: "-1"
  _rest_api_published: "1"
parent_id: "0"
password: ""
status: publish
tags:
- mahout
title: '[Mahout] Deploying custom drivers to mahout'
url: /2016/08/09/mahout-deploying-custom-drivers-to-mahout/
---

Developing custom drivers on Mahout is fairly straightforward. You can inherit
from MahoutDriver for Java drivers and MahourSparkDriver for spark drivers.

The Javadoc for MahoutDriver (if you can find it) provides a good summary of how
to implement it

> [General-purpose driver class for Mahout programs. Utilizes > >
> org.apache.hadoop.util.ProgramDriver to run main methods of other > classes, >
> but first loads up default properties from a properties >
> > file.]{style="color:#353833;font-family:Arial, Helvetica,
> sans-serif;font-size:12.16px;line-height:normal;"}
>
> To run locally:
>
> ```{style="font-size:1.3em;margin-top:0;color:#353833;line-height:normal;"}
> $MAHOUT_HOME/bin/mahout run shortJobName [over-ride ops]
> ```
>
> Works like this: by default, the file "driver.classes.props" is loaded from
> the classpath, which defines a mapping between short names like "vectordump"
> and fully qualified class names. The format of driver.classes.props is like
> so:
>
> ```{style="font-size:1.3em;margin-top:0;color:#353833;line-height:normal;"}
> fully.qualified.class.name = shortJobName : descriptive string
> ```
>
> The default properties to be applied to the program run is pulled out of, by
> default, ".props" (also off of the classpath).
>
> The format of the default properties files is as follows:
>
> ```{style="font-size:1.3em;margin-top:0;color:#353833;line-height:normal;"}
>   i|input = /path/to/my/input
>   o|output = /path/to/my/output
>   m|jarFile = /path/to/jarFile
>   # etc - each line is shortArg|longArg = value
> ```
>
> [The next argument to the Driver is supposed to be the short name of > the >
> class to be run (as defined in the driver.classes.props >
> > file).]{style="color:#353833;font-family:Arial, Helvetica,
> sans-serif;font-size:12.16px;line-height:normal;"}
>
> Then the class which will be run will have it's main called with
>
> ```{style="font-size:1.3em;margin-top:0;color:#353833;line-height:normal;"}
> main(new String[] { "--input", "/path/to/my/input", "--output", "/path/to/my/output" });
> ```
>
> [After all the "default" properties are loaded from the file, any > further
> > command-line arguments are taken in, and over-ride the >
> > defaults.]{style="color:#353833;font-family:Arial, Helvetica,
> sans-serif;font-size:12.16px;line-height:normal;"}
>
> So if your driver.classes.props looks like so:
>
> ```{style="font-size:1.3em;margin-top:0;color:#353833;line-height:normal;"}
> org.apache.mahout.utils.vectors.VectorDumper = vecDump : dump vectors from a sequence file
> ```
>
> [and you have a file core/src/main/resources/vecDump.props which looks >
> > like]{style="color:#353833;font-family:Arial, Helvetica,
> sans-serif;font-size:12.16px;line-height:normal;"}
>
> ```{style="font-size:1.3em;margin-top:0;color:#353833;line-height:normal;"}
>   o|output = /tmp/vectorOut
>   s|seqFile = /my/vector/sequenceFile
> ```
>
> [And you execute the > command-line:]{style="color:#353833;font-family:Arial, Helvetica,
> sans-serif;font-size:12.16px;line-height:normal;"}
>
> ```{style="font-size:1.3em;margin-top:0;color:#353833;line-height:normal;"}
> $MAHOUT_HOME/bin/mahout run vecDump -s /my/otherVector/sequenceFile
> ```
>
> [Then org.apache.mahout.utils.vectors.VectorDumper.main() will be called with
> arguments:

You can also implement it slightly differently by just dumping the jar into the
mahout home directory and naming it starting with "mahout-" i.e.
mahout-mydriver.jar

Hope this helps someone
