Try out Liftbridge with NATS.io ..

Simple demo for Golang KL Aug 2018

Steps: Change Owner name; show changes propagated to SearchListing + Search Service

For simplicity; persistence store in-memory; each service will be a separate go routine with its own sync mechanism.  Shows Value Object method

SearchListing Service
  Sub: Listing_Changes

Listing Service:
  List_Details
  Sub: Owner_Changes, Images_Changes, Car_Changes 

Vehicle Service
  Add_Vehicle
  Delete_Vehicle
  Pub: 

VehicleCatalog Service
  Add_Make
  Delete_Make
  Edit_Make

Image Service
  Add_Image
  Remove_Image
  Pub: Image_Added, Image_Removed

Organization Service
  Sub: User_Changes

User Service
  Pub: Detail_Changes

Moderation Service
  Approve
  Deny
  Pub: Listing_Changes