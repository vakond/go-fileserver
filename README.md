# File server

### Simple file server

This **toy project** is a simple gRPC file server which can be used to organize and distribute files. It explores gRPC stream capabilities.

*    Build both fileserver and client:

         make

*    Usage:

         fileserver start

The first launch creates empty config file with comments. It should be populated with version tags and corresponding filenames.

*    Shutdown:

         Ctrl+C or killall fileserver

*    Simple tests (requires client):

         client versions
         client download
