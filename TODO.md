# TODO

## (DONE!) Upload a photo through the Handler

## Return the image bytes (DONE!)

- Define API return schema of photo metadata (DONE!)
- Develop GET photo/{id}/image endpoint (DONE!)
  - StorageService : connects to GCP to upload photo to cloud
  - PhotoService   : Takes a photo id and looks for it on the DB to get the StorageKey to ask the StorageService to fetch the actual image on cloud
  - PhotoHandler   : Receives the GET request, uses the PhotosService to get the image stream and sens a response to the client

## Add albums

- Define Album model (DONE!)
- Define Album Photo relation (DONE!)
- Create album (DONE!)
  - Album Service createAlbum() (DONE!)
  - Album handler createAlbum() (DONE!)
- Add photo to album (DONE!)
  - Album Service addPhotos() (DONE!)
  - Album handler addPhotos() (DONE!)
- Get all albums (DONE!)
  - Album Service getAlbums() (DONE!)
  - Album handler getAlbums() (DONE!)
- Get album by id (DONE!)
  - Album Service getAlbum() (DONE!)
  - Album handler getAlbum() (DONE!)

## Return Signed URL when returning photo image instead of bytes (In Progress... 🚧) (DONE!)

## Refactor schemas (DONE!)

- merge all schemas spread in different files into one for each resource (DONE!)
- Delete unused schemas (DONE!)

## Update album, SET cover image (DONE!)

## Upload to Artifact registry

## Back-end MVP IS NOW COMPLETE !!!

## Return the re-seized image

## Return re-sized and watermarked image

## Make a cart

## Shoping cart CRUD

## Cart checkout

## Return Full image if bought
