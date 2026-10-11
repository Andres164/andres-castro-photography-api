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
- Create album
  - Album Service createAlbum()
  - Album handler createAlbum()
- Add photo to album
  - Album Service addPhotos()
  - Album handler addPhotos()
- Get all albums
  - Album Service getAlbums()
  - Album handler getAlbums()
- Get album by id
  - Album Service getAlbum()
  - Album handler getAlbum()

## Clean Schemas delete the ones that aren't used anymore

## Return the re-seized image

## Return re-sized and watermarked image

## Make a cart

## Shoping cart CRUD

## Cart checkout

## Return Full image if bought
