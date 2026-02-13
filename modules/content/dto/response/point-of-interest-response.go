package responseContentDTO

type CategoryCode string

const (
	CategoryCodeAirport                                  CategoryCode = "AIRPORT"
	CategoryCodeAmusementPark                            CategoryCode = "AMUSEMENT_PARK"
	CategoryCodeAquarium                                 CategoryCode = "AQUARIUM"
	CategoryCodeBeach                                    CategoryCode = "BEACH"
	CategoryCodeBoatDock                                 CategoryCode = "BOAT_DOCK"
	CategoryCodeBusStation                               CategoryCode = "BUS_STATION"
	CategoryCodeBusinessLocation                         CategoryCode = "BUSINESS_LOCATION"
	CategoryCodeCanal                                    CategoryCode = "CANAL"
	CategoryCodeCarRentalLocation                        CategoryCode = "CAR_RENTAL_LOCATION"
	CategoryCodeCasino                                   CategoryCode = "CASINO"
	CategoryCodeCemetery                                 CategoryCode = "CEMETERY"
	CategoryCodeChurch                                   CategoryCode = "CHURCH"
	CategoryCodeConcertHall                              CategoryCode = "CONCERT_HALL"
	CategoryCodeConferenceCenter                         CategoryCode = "CONFERENCE_CENTER"
	CategoryCodeConventionCenter                         CategoryCode = "CONVENTION_CENTER"
	CategoryCodeFairground                               CategoryCode = "FAIRGROUND"
	CategoryCodeFarm                                     CategoryCode = "FARM"
	CategoryCodeGallery                                  CategoryCode = "GALLERY"
	CategoryCodeHistoricBuilding                         CategoryCode = "HISTORIC_BUILDING"
	CategoryCodeHospital                                 CategoryCode = "HOSPITAL"
	CategoryCodeLake                                     CategoryCode = "LAKE"
	CategoryCodeLandmark                                 CategoryCode = "LANDMARK"
	CategoryCodeMarina                                   CategoryCode = "MARINA"
	CategoryCodeMarket                                   CategoryCode = "MARKET"
	CategoryCodeMonument                                 CategoryCode = "MONUMENT"
	CategoryCodeMountain                                 CategoryCode = "MOUNTAIN"
	CategoryCodeMuseum                                   CategoryCode = "MUSEUM"
	CategoryCodeOcean                                    CategoryCode = "OCEAN"
	CategoryCodePalace                                   CategoryCode = "PALACE"
	CategoryCodePark                                     CategoryCode = "PARK"
	CategoryCodeRecreationCenter                         CategoryCode = "RECREATION_CENTER"
	CategoryCodeRestaurant                               CategoryCode = "RESTAURANT"
	CategoryCodeRiver                                    CategoryCode = "RIVER"
	CategoryCodeShoppingMall                             CategoryCode = "SHOPPING_MALL"
	CategoryCodeSkiArea                                  CategoryCode = "SKI_AREA"
	CategoryCodeStadium                                  CategoryCode = "STADIUM"
	CategoryCodeStore                                    CategoryCode = "STORE"
	CategoryCodeTheaterCinema                            CategoryCode = "THEATER_CINEMA"
	CategoryCodeTrainStation                             CategoryCode = "TRAIN_STATION"
	CategoryCodeUniversity                               CategoryCode = "UNIVERSITY"
	CategoryCodeWinery                                   CategoryCode = "WINERY"
	CategoryCodeZoo                                      CategoryCode = "ZOO"
	CategoryCodeCityEvent                                CategoryCode = "CITY_EVENT"
	CategoryCodeFestival                                 CategoryCode = "FESTIVAL"
	CategoryCodeTour                                     CategoryCode = "TOUR"
	CategoryCodeOther                                    CategoryCode = "OTHER"
	CategoryCodeNightlife                                CategoryCode = "NIGHTLIFE"
	CategoryCodeShopping                                 CategoryCode = "SHOPPING"
	CategoryCodeSports                                   CategoryCode = "SPORTS"
	CategoryCodeCityCenter                               CategoryCode = "CITY_CENTER"
	CategoryCodeCityDowntown                             CategoryCode = "CITY_DOWNTOWN"
	CategoryCodeLivetheater                              CategoryCode = "LIVETHEATER"
	CategoryCodeArena                                    CategoryCode = "ARENA"
	CategoryCodeBar                                      CategoryCode = "BAR"
	CategoryCodeBay                                      CategoryCode = "BAY"
	CategoryCodeCathedral                                CategoryCode = "CATHEDRAL"
	CategoryCodeEducationalInstitution                   CategoryCode = "EDUCATIONAL_INSTITUTION"
	CategoryCodeMedicalFacility                          CategoryCode = "MEDICAL_FACILITY"
	CategoryCodeArmyBase                                 CategoryCode = "ARMY_BASE"
	CategoryCodeCommercialDistrict                       CategoryCode = "COMMERCIAL_DISTRICT"
	CategoryCodeTouristSite                              CategoryCode = "TOURIST_SITE"
	CategoryCodeMiscellaneous                            CategoryCode = "MISCELLANEOUS"
	CategoryCodeAgricultural                             CategoryCode = "AGRICULTURAL"
	CategoryCodeArcheological                            CategoryCode = "ARCHEOLOGICAL"
	CategoryCodeBotanicalGarden                          CategoryCode = "BOTANICAL_GARDEN"
	CategoryCodeBowling                                  CategoryCode = "BOWLING"
	CategoryCodeCulturalCenter                           CategoryCode = "CULTURAL_CENTER"
	CategoryCodeEquestrianCenter                         CategoryCode = "EQUESTRIAN_CENTER"
	CategoryCodeHandicraftCenter                         CategoryCode = "HANDICRAFT_CENTER"
	CategoryCodeNaturalAttraction                        CategoryCode = "NATURAL_ATTRACTION"
	CategoryCodePerformingArtCenter                      CategoryCode = "PERFORMING_ART_CENTER"
	CategoryCodePlanetariumScienceCenter                 CategoryCode = "PLANETARIUM_SCIENCE_CENTER"
	CategoryCodeCableCars                                CategoryCode = "CABLE_CARS"
	CategoryCodeCompany                                  CategoryCode = "COMPANY"
	CategoryCodeFactoryBusinessTour                      CategoryCode = "FACTORY_BUSINESS_TOUR"
	CategoryCodeNighttimeEntertainment                   CategoryCode = "NIGHTTIME_ENTERTAINMENT"
	CategoryCodeArt                                      CategoryCode = "ART"
	CategoryCodeMusic                                    CategoryCode = "MUSIC"
	CategoryCodeStateNationalPark                        CategoryCode = "STATE_NATIONAL_PARK"
	CategoryCodeExhibitionConferenceCenter               CategoryCode = "EXHIBITION_CONFERENCE_CENTER"
	CategoryCodeAirLineDesk                              CategoryCode = "AIR_LINE_DESK"
	CategoryCodeAnimalWatching                           CategoryCode = "ANIMAL_WATCHING"
	CategoryCodeAtmCashMachine                           CategoryCode = "ATM_CASH_MACHINE"
	CategoryCodeBabySitting                              CategoryCode = "BABY_SITTING"
	CategoryCodeBaggageStorage                           CategoryCode = "BAGGAGE_STORAGE"
	CategoryCodeBallroom                                 CategoryCode = "BALLROOM"
	CategoryCodeBeachNearHotel                           CategoryCode = "BEACH_NEAR_HOTEL"
	CategoryCodeHotelWithDirectAccessToABeach            CategoryCode = "HOTEL_WITH_DIRECT_ACCESS_TO_A_BEACH"
	CategoryCodeBirdWatching                             CategoryCode = "BIRD_WATCHING"
	CategoryCodeRoomsWithBalcony                         CategoryCode = "ROOMS_WITH_BALCONY"
	CategoryCodeBoating                                  CategoryCode = "BOATING"
	CategoryCodeBeautyParlour                            CategoryCode = "BEAUTY_PARLOUR"
	CategoryCodeCoachBusParking                          CategoryCode = "COACH_BUS_PARKING"
	CategoryCodeButlerService                            CategoryCode = "BUTLER_SERVICE"
	CategoryCodeCarRental                                CategoryCode = "CAR_RENTAL"
	CategoryCodeChildrenWelcome                          CategoryCode = "CHILDREN_WELCOME"
	CategoryCodeChildrenNotAllowed                       CategoryCode = "CHILDREN_NOT_ALLOWED"
	CategoryCodeConnectingRooms                          CategoryCode = "CONNECTING_ROOMS"
	CategoryCodeConcierge                                CategoryCode = "CONCIERGE"
	CategoryCodeCourtesyCar                              CategoryCode = "COURTESY_CAR"
	CategoryCodeCellularPhoneRental                      CategoryCode = "CELLULAR_PHONE_RENTAL"
	CategoryCodeDutyFreeShop                             CategoryCode = "DUTY_FREE_SHOP"
	CategoryCodeDisco                                    CategoryCode = "DISCO"
	CategoryCodeDrivingRange                             CategoryCode = "DRIVING_RANGE"
	CategoryCodeElevator                                 CategoryCode = "ELEVATOR"
	CategoryCodeLiveEntertainment                        CategoryCode = "LIVE_ENTERTAINMENT"
	CategoryCodeCurrencyExchangeFacilities               CategoryCode = "CURRENCY_EXCHANGE_FACILITIES"
	CategoryCodeExecutiveDesk                            CategoryCode = "EXECUTIVE_DESK"
	CategoryCodeExecutiveFloor                           CategoryCode = "EXECUTIVE_FLOOR"
	CategoryCodeExpressCheckIn                           CategoryCode = "EXPRESS_CHECK_IN"
	CategoryCodeExpressCheckOut                          CategoryCode = "EXPRESS_CHECK_OUT"
	CategoryCodeFrontDeskOpen24HoursADay                 CategoryCode = "FRONT_DESK_OPEN_24_HOURS_A_DAY"
	CategoryCodeFishing                                  CategoryCode = "FISHING"
	CategoryCodeFlorist                                  CategoryCode = "FLORIST"
	CategoryCodeFreeParking                              CategoryCode = "FREE_PARKING"
	CategoryCodeFreeTransportation                       CategoryCode = "FREE_TRANSPORTATION"
	CategoryCodeGamesRoom                                CategoryCode = "GAMES_ROOM"
	CategoryCodeGarageParking                            CategoryCode = "GARAGE_PARKING"
	CategoryCodeGiftShopNewsStand                        CategoryCode = "GIFT_SHOP_NEWS_STAND"
	CategoryCodeGolf                                     CategoryCode = "GOLF"
	CategoryCodeGymNotHealthClub                         CategoryCode = "GYM_NOT_HEALTH_CLUB"
	CategoryCodeHealthClub                               CategoryCode = "HEALTH_CLUB"
	CategoryCodeHorseRiding                              CategoryCode = "HORSE_RIDING"
	CategoryCodeHotspots                                 CategoryCode = "HOTSPOTS"
	CategoryCodeFreeHighSpeedInternetConnection          CategoryCode = "FREE_HIGH_SPEED_INTERNET_CONNECTION"
	CategoryCodeHighSpeedInternetConnection              CategoryCode = "HIGH_SPEED_INTERNET_CONNECTION"
	CategoryCodeInternetServices                         CategoryCode = "INTERNET_SERVICES"
	CategoryCodeJacuzzi                                  CategoryCode = "JACUZZI"
	CategoryCodeJoggingTrack                             CategoryCode = "JOGGING_TRACK"
	CategoryCodeKennels                                  CategoryCode = "KENNELS"
	CategoryCodeLaundryService                           CategoryCode = "LAUNDRY_SERVICE"
	CategoryCodeMassage                                  CategoryCode = "MASSAGE"
	CategoryCodeMiniatureGolf                            CategoryCode = "MINIATURE_GOLF"
	CategoryCodeMultilingualStaff                        CategoryCode = "MULTILINGUAL_STAFF"
	CategoryCodeNightClub                                CategoryCode = "NIGHT_CLUB"
	CategoryCodeHotelDoesNotProvidePornographicFilmsTv   CategoryCode = "HOTEL_DOES_NOT_PROVIDE_PORNOGRAPHIC_FILMS_TV"
	CategoryCodeNursery                                  CategoryCode = "NURSERY"
	CategoryCodeParking                                  CategoryCode = "PARKING"
	CategoryCodePetsAllowed                              CategoryCode = "PETS_ALLOWED"
	CategoryCodePharmacy                                 CategoryCode = "PHARMACY"
	CategoryCodeChildrenSPlayArea                        CategoryCode = "CHILDREN_S_PLAY_AREA"
	CategoryCodePorterBellBoy                            CategoryCode = "PORTER_BELL_BOY"
	CategoryCodePuttingGreen                             CategoryCode = "PUTTING_GREEN"
	CategoryCodeSauna                                    CategoryCode = "SAUNA"
	CategoryCodeScubaDiving                              CategoryCode = "SCUBA_DIVING"
	CategoryCodeFreeAirportShuttle                       CategoryCode = "FREE_AIRPORT_SHUTTLE"
	CategoryCodeIndoorSwimmingPool                       CategoryCode = "INDOOR_SWIMMING_POOL"
	CategoryCodeSightseeing                              CategoryCode = "SIGHTSEEING"
	CategoryCodeSkeetShooting                            CategoryCode = "SKEET_SHOOTING"
	CategoryCodeHotelWithSkiInOutFacilities              CategoryCode = "HOTEL_WITH_SKI_IN_OUT_FACILITIES"
	CategoryCodeSnowSkiing                               CategoryCode = "SNOW_SKIING"
	CategoryCodeSolarium                                 CategoryCode = "SOLARIUM"
	CategoryCodeSpa                                      CategoryCode = "SPA"
	CategoryCodeHeatedSwimmingPool                       CategoryCode = "HEATED_SWIMMING_POOL"
	CategoryCodeSwimmingPool                             CategoryCode = "SWIMMING_POOL"
	CategoryCodeIndoorTennis                             CategoryCode = "INDOOR_TENNIS"
	CategoryCodeTennis                                   CategoryCode = "TENNIS"
	CategoryCodeTennisProfessional                       CategoryCode = "TENNIS_PROFESSIONAL"
	CategoryCodeTheatreDesk                              CategoryCode = "THEATRE_DESK"
	CategoryCodeTourDesk                                 CategoryCode = "TOUR_DESK"
	CategoryCodeTranslationServices                      CategoryCode = "TRANSLATION_SERVICES"
	CategoryCodeTravelAgency                             CategoryCode = "TRAVEL_AGENCY"
	CategoryCodeValetParking                             CategoryCode = "VALET_PARKING"
	CategoryCodeVendingMachines                          CategoryCode = "VENDING_MACHINES"
	CategoryCodeVolleyball                               CategoryCode = "VOLLEYBALL"
	CategoryCodeWaterSports                              CategoryCode = "WATER_SPORTS"
	CategoryCodeWirelessConnectivity                     CategoryCode = "WIRELESS_CONNECTIVITY"
	CategoryCodeWeddingServices                          CategoryCode = "WEDDING_SERVICES"
	CategoryCodeHairDresser                              CategoryCode = "HAIR_DRESSER"
	CategoryCodeBusinessServices                         CategoryCode = "BUSINESS_SERVICES"
	CategoryCodeAccessibleFacilities                     CategoryCode = "ACCESSIBLE_FACILITIES"
	CategoryCodeSecurity                                 CategoryCode = "SECURITY"
	CategoryCodeGroupRates                               CategoryCode = "GROUP_RATES"
	CategoryCode24HourSecurity                           CategoryCode = "24_HOUR_SECURITY"
	CategoryCodePhotocopyCenter                          CategoryCode = "PHOTOCOPY_CENTER"
	CategoryCodeVideoTapes                               CategoryCode = "VIDEO_TAPES"
	CategoryCodeWakeupService                            CategoryCode = "WAKEUP_SERVICE"
	CategoryCodeDirectDialTelephone                      CategoryCode = "DIRECT_DIAL_TELEPHONE"
	CategoryCodeEarlyCheckIn                             CategoryCode = "EARLY_CHECK_IN"
	CategoryCodeBicycleRentals                           CategoryCode = "BICYCLE_RENTALS"
	CategoryCodeLateCheckOutAvailable                    CategoryCode = "LATE_CHECK_OUT_AVAILABLE"
	CategoryCodeBookstore                                CategoryCode = "BOOKSTORE"
	CategoryCodeComplimentarySelfServiceLaundry          CategoryCode = "COMPLIMENTARY_SELF_SERVICE_LAUNDRY"
	CategoryCodeAccessibleParking                        CategoryCode = "ACCESSIBLE_PARKING"
	CategoryCodeBoutiquesStores                          CategoryCode = "BOUTIQUES_STORES"
	CategoryCodeShopsAndCommercialServices               CategoryCode = "SHOPS_AND_COMMERCIAL_SERVICES"
	CategoryCodeSportsBarOpenForLunch                    CategoryCode = "SPORTS_BAR_OPEN_FOR_LUNCH"
	CategoryCodeComplimentaryCoffeeInLobby               CategoryCode = "COMPLIMENTARY_COFFEE_IN_LOBBY"
	CategoryCodeDinnerDeliveryServiceFromLocalRestaurant CategoryCode = "DINNER_DELIVERY_SERVICE_FROM_LOCAL_RESTAURANT"
	CategoryCodeComplimentaryNewspaperInLobby            CategoryCode = "COMPLIMENARY_NEWSPAPER_IN_LOBBY"
	CategoryCodeFrontDesk                                CategoryCode = "FRONT_DESK"
	CategoryCodeGroceryShoppingServiceAvailable          CategoryCode = "GROCERY_SHOPPING_SERVICE_AVAILABLE"
	CategoryCodeManagersReception                        CategoryCode = "MANAGERS_RECEPTION"
	CategoryCodeMedicalFacilitiesService                 CategoryCode = "MEDICAL_FACILITIES_SERVICE"
	CategoryCodeAllInclusiveMealPlan                     CategoryCode = "ALL_INCLUSIVE_MEAL_PLAN"
	CategoryCodeCommunalBarArea                          CategoryCode = "COMMUNAL_BAR_AREA"
	CategoryCodeContinentalBreakfast                     CategoryCode = "CONTINENTAL_BREAKFAST"
	CategoryCodeFullMealPlan                             CategoryCode = "FULL_MEAL_PLAN"
	CategoryCodeOnsiteLaundry                            CategoryCode = "ONSITE_LAUNDRY"
	CategoryCode24HourFoodBeverageKiosk                  CategoryCode = "24_HOUR_FOOD_BEVERAGE_KIOSK"
	CategoryCodeFullServiceHousekeeping                  CategoryCode = "FULL_SERVICE_HOUSEKEEPING"
	CategoryCodeAdditionalServicesAmenitiesFacilities    CategoryCode = "ADDITIONAL_SERVICES_AMENITIES_FACILITIES_ON_PROPERTY"
	CategoryCodeDvdVideoRental                           CategoryCode = "DVD_VIDEO_RENTAL"
	CategoryCodeParkingLot                               CategoryCode = "PARKING_LOT"
	CategoryCodeCocktailLoungeWithEntertainment          CategoryCode = "COCKTAIL_LOUNGE_WITH_ENTERTAINMENT"
	CategoryCodeCocktailLounge                           CategoryCode = "COCKTAIL_LOUNGE"
	CategoryCodePhoneServices                            CategoryCode = "PHONE_SERVICES"
	CategoryCodeAerobicsInstruction                      CategoryCode = "AEROBICS_INSTRUCTION"
	CategoryCodeCoinOperatedLaundry                      CategoryCode = "COIN_OPERATED_LAUNDRY"
	CategoryCodeBankingServices                          CategoryCode = "BANKING_SERVICES"
	CategoryCodeExhibitionConventionFloor                CategoryCode = "EXHIBITION_CONVENTION_FLOOR"
	CategoryCodeCourtyard                                CategoryCode = "COURTYARD"
	CategoryCodeDoorMan                                  CategoryCode = "DOOR_MAN"
	CategoryCodeDrugstorePharmacy                        CategoryCode = "DRUGSTORE_PHARMACY"
	CategoryCodeHousekeepingDaily                        CategoryCode = "HOUSEKEEPING_DAILY"
	CategoryCodeOffSiteParking                           CategoryCode = "OFF_SITE_PARKING"
	CategoryCodeOnSiteParking                            CategoryCode = "ON_SITE_PARKING"
	CategoryCodeOutdoorParking                           CategoryCode = "OUTDOOR_PARKING"
	CategoryCodeRampAccess                               CategoryCode = "RAMP_ACCESS"
	CategoryCodeSportsBar                                CategoryCode = "SPORTS_BAR"
	CategoryCodeValetDryCleaning                         CategoryCode = "VALET_DRY_CLEANING"
	CategoryCodeChildrensProgramOnsite                   CategoryCode = "CHILDRENS_PROGRAM_ONSITE"
	CategoryCodeWindsurfing                              CategoryCode = "WINDSURFING"
	CategoryCodeCamping                                  CategoryCode = "CAMPING"
	CategoryCodeHunting                                  CategoryCode = "HUNTING"
	CategoryCodeIndoorOutdoorConnectingPool              CategoryCode = "INDOOR_OUTDOOR_CONNECTING_POOL"
	CategoryCodeMountainClimbing                         CategoryCode = "MOUNTAIN_CLIMBING"
	CategoryCodeNaturePreserveTrail                      CategoryCode = "NATURE_PRESERVE_TRAIL"
	CategoryCodeBilliards                                CategoryCode = "BILLIARDS"
	CategoryCodeSunTanningBed                            CategoryCode = "SUN_TANNING_BED"
	CategoryCodeSurfing                                  CategoryCode = "SURFING"
	CategoryCodeTableTennis                              CategoryCode = "TABLE_TENNIS"
	CategoryCodeTeenPrograms                             CategoryCode = "TEEN_PROGRAMS"
	CategoryCodeIndoorPool                               CategoryCode = "INDOOR_POOL"
	CategoryCodeOutdoorPool                              CategoryCode = "OUTDOOR_POOL"
	CategoryCodeChildrensProgram                         CategoryCode = "CHILDRENS_PROGRAM"
	CategoryCodeBoxing                                   CategoryCode = "BOXING"
	CategoryCodeChildrensPool                            CategoryCode = "CHILDRENS_POOL"
	CategoryCodeDancing                                  CategoryCode = "DANCING"
	CategoryCodeGarden                                   CategoryCode = "GARDEN"
	CategoryCodeKaraoke                                  CategoryCode = "KARAOKE"
	CategoryCodeMuseumGalleryViewing                     CategoryCode = "MUSEUM_GALLERY_VIEWING"
	CategoryCodeNightclubs                               CategoryCode = "NIGHTCLUBS"
	CategoryCodeSportsEvents                             CategoryCode = "SPORTS_EVENTS"
	CategoryCodeSkydiving                                CategoryCode = "SKYDIVING"
	CategoryCodeSunbathing                               CategoryCode = "SUNBATHING"
	CategoryCodeTheatre                                  CategoryCode = "THEATRE"
	CategoryCodeFitnessCenterOffSite                     CategoryCode = "FITNESS_CENTER_OFF_SITE"
	CategoryCodeFlyFishing                               CategoryCode = "FLY_FISHING"
	CategoryCodeBaseballDiamond                          CategoryCode = "BASEBALL_DIAMOND"
	CategoryCodeGym                                      CategoryCode = "GYM"
	CategoryCodeBasketballCourt                          CategoryCode = "BASKETBALL_COURT"
	CategoryCodeBikeTrail                                CategoryCode = "BIKE_TRAIL"
	CategoryCodeHikingTrail                              CategoryCode = "HIKING_TRAIL"
	CategoryCodeJoggingTrail                             CategoryCode = "JOGGING_TRAIL"
	CategoryCodeKayaking                                 CategoryCode = "KAYAKING"
	CategoryCodeMountainBikingTrail                      CategoryCode = "MOUNTAIN_BIKING_TRAIL"
	CategoryCodeParasailing                              CategoryCode = "PARASAILING"
	CategoryCodePlayground                               CategoryCode = "PLAYGROUND"
	CategoryCodePool                                     CategoryCode = "POOL"
	CategoryCodeRiverRafting                             CategoryCode = "RIVER_RAFTING"
	CategoryCodeSailing                                  CategoryCode = "SAILING"
	CategoryCodeSnorkeling                               CategoryCode = "SNORKELING"
	CategoryCodeTennisCourt                              CategoryCode = "TENNIS_COURT"
	CategoryCodeWaterSkiing                              CategoryCode = "WATER_SKIING"
	CategoryCodeFineDining                               CategoryCode = "FINE_DINING"
	CategoryCodeGolfLocation                             CategoryCode = "GOLF_LOCATION"
	CategoryCodeBilingualStaff                           CategoryCode = "BILINGUAL_STAFF"
	CategoryCodeAirConditioning                          CategoryCode = "AIR_CONDITIONING"
	CategoryCodeNonSmokingRooms                          CategoryCode = "NON_SMOKING_ROOMS"
	CategoryCodeInternetAccess                           CategoryCode = "INTERNET_ACCESS"
	CategoryCodeSundryConvenienceStore                   CategoryCode = "SUNDRY_CONVENIENCE_STORE"
	CategoryCodeTransportation                           CategoryCode = "TRANSPORTATION"
	CategoryCodeComplimentaryBreakfast                   CategoryCode = "COMPLIMENTARY_BREAKFAST"
	CategoryCodeHighSpeedInternetAccess                  CategoryCode = "HIGH_SPEED_INTERNET_ACCESS"
	CategoryCodeLobby                                    CategoryCode = "LOBBY"
	CategoryCode24HourCoffeeShop                         CategoryCode = "24_HOUR_COFFEE_SHOP"
	CategoryCodeAirportShuttleService                    CategoryCode = "AIRPORT_SHUTTLE_SERVICE"
	CategoryCodeLuggageService                           CategoryCode = "LUGGAGE_SERVICE"
	CategoryCodePianoBar                                 CategoryCode = "PIANO_BAR"
	CategoryCodeVipSecurity                              CategoryCode = "VIP_SECURITY"
	CategoryCodeWheelChairAccess                         CategoryCode = "WHEEL_CHAIR_ACCESS"
	CategoryCodeBusinessCenter                           CategoryCode = "BUSINESS_CENTER"
	CategoryCodeChildPrograms                            CategoryCode = "CHILD_PROGRAMS"
	CategoryCodeSeaside                                  CategoryCode = "SEASIDE"
	CategoryCodePrivateDiningForGroups                   CategoryCode = "PRIVATE_DINING_FOR_GROUPS"
	CategoryCodeHighSpeedWireless                        CategoryCode = "HIGH_SPEED_WIRELESS"
	CategoryCodePrinter                                  CategoryCode = "PRINTER"
	CategoryCodeIfGuestRoomsHaveMoreThanOnePhoneLine     CategoryCode = "IF_GUEST_ROOMS_HAVE_MORE_THAN_ONE_PHONE_LINE"
	CategoryCodeComplimentaryWirelessInternet            CategoryCode = "COMPLIMENTARY_WIRELESS_INTERNET"
	CategoryCodeSameGenderFloor                          CategoryCode = "SAME_GENDER_FLOOR"
	CategoryCodeChildrenPrograms                         CategoryCode = "CHILDREN_PROGRAMS"
	CategoryCodeBuildingMeetsLocal                       CategoryCode = "BUILDING_MEETS_LOCAL"
	CategoryCodeInternetBrowserOnTv                      CategoryCode = "INTERNET_BROWSER_ON_TV"
	CategoryCodeNewspaper                                CategoryCode = "NEWSPAPER"
	CategoryCodeParkingControlledAccessGates             CategoryCode = "PARKING_CONTROLLED_ACCESS_GATES_TO_ENTER_PARKING_AREA"
	CategoryCodeHotelSafeDepositBoxNotRoomSafeBox        CategoryCode = "HOTEL_SAFE_DEPOSIT_BOX__NOT_ROOM_SAFE_BOX"
	CategoryCodeStorageSpaceAvailableForFee              CategoryCode = "STORAGE_SPACE_AVAILABLE_FOR_FEE"
	CategoryCodeTypeOfEntranceToGuestRoom                CategoryCode = "TYPE_OF_ENTRANCE_TO_GUEST_ROOM"
	CategoryCodeBeverageCocktail                         CategoryCode = "BEVERAGE_COCKTAIL"
	CategoryCodeCellPhoneRental                          CategoryCode = "CELL_PHONE_RENTAL"
	CategoryCodeCoffeeTea                                CategoryCode = "COFFEE_TEA"
	CategoryCodeEarlyCheckInGuarantee                    CategoryCode = "EARLY_CHECK_IN_GUARANTEE"
	CategoryCodeFoodAndBeverageDiscount                  CategoryCode = "FOOD_AND_BEVERAGE_DISCOUNT"
	CategoryCodeLateCheckOutGuarantee                    CategoryCode = "LATE_CHECK_OUT_GUARANTEE"
	CategoryCodeRoomUpgradeConfirmed                     CategoryCode = "ROOM_UPGRADE_CONFIRMED"
	CategoryCodeRoomUpgradeOnAvailability                CategoryCode = "ROOM_UPGRADE_ON_AVAILABILITY"
	CategoryCodeShuttleToLocalBusinesses                 CategoryCode = "SHUTTLE_TO_LOCAL_BUSINESSES"
	CategoryCodeShuttleToLocalAttractions                CategoryCode = "SHUTTLE_TO_LOCAL_ATTRACTIONS"
	CategoryCodeSocialHour                               CategoryCode = "SOCIAL_HOUR"
	CategoryCodeVideoBilling                             CategoryCode = "VIDEO_BILLING"
	CategoryCodeWelcomeGift                              CategoryCode = "WELCOME_GIFT"
	CategoryCodeHypoallergenicRooms                      CategoryCode = "HYPOALLERGENIC_ROOMS"
	CategoryCodeRoomAirFiltration                        CategoryCode = "ROOM_AIR_FILTRATION"
	CategoryCodeSmokeFreeProperty                        CategoryCode = "SMOKE_FREE_PROPERTY"
	CategoryCodeWaterPurificationSystemInUse             CategoryCode = "WATER_PURIFICATION_SYSTEM_IN_USE"
	CategoryCodePoolsideService                          CategoryCode = "POOLSIDE_SERVICE"
	CategoryCodeClothingStore                            CategoryCode = "CLOTHING_STORE"
	CategoryCodeEvElectricVehicleChargingLocation        CategoryCode = "EV_ELECTRIC_VEHICLE_CHARGING_LOCATION"
	CategoryCodeOfficeRental                             CategoryCode = "OFFICE_RENTAL"
	CategoryCodeIncomingFax                              CategoryCode = "INCOMING_FAX"
	CategoryCodeOutgoingFax                              CategoryCode = "OUTGOING_FAX"
	CategoryCodeBabyKit                                  CategoryCode = "BABY_KIT"
	CategoryCodeChildrenSBreakfast                       CategoryCode = "CHILDREN_S_BREAKFAST"
	CategoryCodeCloakroomService                         CategoryCode = "CLOAKROOM_SERVICE"
	CategoryCodeCoffeeLounge                             CategoryCode = "COFFEE_LOUNGE"
	CategoryCodeEventsTicketService                      CategoryCode = "EVENTS_TICKET_SERVICE"
	CategoryCodeLateCheckIn                              CategoryCode = "LATE_CHECK_IN"
	CategoryCodeLimitedParking                           CategoryCode = "LIMITED_PARKING"
	CategoryCodeOutdoorSummerBarCafe                     CategoryCode = "OUTDOOR_SUMMER_BAR_CAFE"
	CategoryCodeNoParkingAvailable                       CategoryCode = "NO_PARKING_AVAILABLE"
	CategoryCodeBeerGarden                               CategoryCode = "BEER_GARDEN"
	CategoryCodeGardenLoungeBar                          CategoryCode = "GARDEN_LOUNGE_BAR"
	CategoryCodeSummerTerrace                            CategoryCode = "SUMMER_TERRACE"
	CategoryCodeWinterTerrace                            CategoryCode = "WINTER_TERRACE"
	CategoryCodeRoofTerrace                              CategoryCode = "ROOF_TERRACE"
	CategoryCodeBeachBar                                 CategoryCode = "BEACH_BAR"
	CategoryCodeHelicopterService                        CategoryCode = "HELICOPTER_SERVICE"
	CategoryCodeFerry                                    CategoryCode = "FERRY"
	CategoryCodeTapasBar                                 CategoryCode = "TAPAS_BAR"
	CategoryCodeCafeBar                                  CategoryCode = "CAFE_BAR"
	CategoryCodeSnackBar                                 CategoryCode = "SNACK_BAR"
	CategoryCodeEnhancedSafetyProtocol                   CategoryCode = "ENHANCED_SAFETY_PROTOCOL"
	CategoryCodeBusinessLibrary                          CategoryCode = "BUSINESS_LIBRARY"
	CategoryCodeCheckInKioskAvailable                    CategoryCode = "CHECK_IN_KIOSK_AVAILABLE"
	CategoryCodeConciergeFloor                           CategoryCode = "CONCIERGE_FLOOR"
	CategoryCodeHousekeepingWeekly                       CategoryCode = "HOUSEKEEPING_WEEKLY"
	CategoryCodePackageReceiving                         CategoryCode = "PACKAGE_RECEIVING"
	CategoryCodePublicAddressSystem                      CategoryCode = "PUBLIC_ADDRESS_SYSTEM"
	CategoryCodeShoeShine                                CategoryCode = "SHOE_SHINE"
	CategoryCodeStorageSpaceAvailable                    CategoryCode = "STORAGE_SPACE_AVAILABLE"
	CategoryCodeTechnicalConciergeAvailable              CategoryCode = "TECHNICAL_CONCIERGE_AVAILABLE"
	CategoryCodeTruckParking                             CategoryCode = "TRUCK_PARKING"
	CategoryCodeWakeUpCalls                              CategoryCode = "WAKE_UP_CALLS"
	CategoryCodeVideoGames                               CategoryCode = "VIDEO_GAMES"
	CategoryCodeRoomServiceLimitedHours                  CategoryCode = "ROOM_SERVICE_LIMITED_HOURS"
	CategoryCodePublicAreasAirConditioned                CategoryCode = "PUBLIC_AREAS_AIR_CONDITIONED"
	CategoryCodeComplimentaryInRoomCoffeeOrTea           CategoryCode = "COMPLIMENTARY_IN_ROOM_COFFEE_OR_TEA"
	CategoryCodeComplimentaryBuffetBreakfast             CategoryCode = "COMPLIMENTARY_BUFFET_BREAKFAST"
	CategoryCodeComplimentaryContinentalBreakfast        CategoryCode = "COMPLIMENTARY_CONTINENTAL_BREAKFAST"
	CategoryCodeLimousineService                         CategoryCode = "LIMOUSINE_SERVICE"
	CategoryCodeTelephoneJackAdaptorAvailable            CategoryCode = "TELEPHONE_JACK_ADAPTOR_AVAILABLE"
	CategoryCodeBreakfastFull                            CategoryCode = "BREAKFAST_FULL"
	CategoryCodeVipLounge                                CategoryCode = "VIP_LOUNGE"
	CategoryCodeParkingFeeManagedByTheHotel              CategoryCode = "PARKING_FEE_MANAGED_BY_THE_HOTEL"
	CategoryCodeHousekeepingLimited                      CategoryCode = "HOUSEKEEPING_LIMITED"
	CategoryCodeTransportationServicesLocalArea          CategoryCode = "TRANSPORTATION_SERVICES_LOCAL_AREA"
	CategoryCodeTransportationServicesLocalOffice        CategoryCode = "TRANSPORTATION_SERVICES_LOCAL_OFFICE"
	CategoryCodeParkingDeck                              CategoryCode = "PARKING_DECK"
	CategoryCodeParkingSideStreet                        CategoryCode = "PARKING_SIDE_STREET"
	CategoryCodeCocktailLoungeWithLightFare              CategoryCode = "COCKTAIL_LOUNGE_WITH_LIGHT_FARE"
	CategoryCodeMotorcycleParking                        CategoryCode = "MOTORCYCLE_PARKING"
	CategoryCodePersonalTrainer                          CategoryCode = "PERSONAL_TRAINER"
	CategoryCodeJetskiing                                CategoryCode = "JETSKIING"
	CategoryCodeRacquetballCourt                         CategoryCode = "RACQUETBALLCOURT"
	CategoryCodeSquashCourts                             CategoryCode = "SQUASHCOURTS"
	CategoryCodeSteamBath                                CategoryCode = "STEAM_BATH"
	CategoryCodeWhirlpool                                CategoryCode = "WHIRLPOOL"
	CategoryCodeSafari                                   CategoryCode = "SAFARI"
	CategoryCodeRecreationSportsCourt                    CategoryCode = "RECREATION_SPORTS_COURT"
	CategoryCodeSnowmobiling                             CategoryCode = "SNOWMOBILING"
	CategoryCodePolo                                     CategoryCode = "POLO"
	CategoryCodeWeightliftingEquipment                   CategoryCode = "WEIGHTLIFTINGEQUIPMENT"
	CategoryCodeCardiovascularEquipment                  CategoryCode = "CARDIOVASCULAREQUIPMENT"
	CategoryCodeExtensiveHealthClub                      CategoryCode = "EXTENSIVEHEALTHCLUB"
	CategoryCodeLimitedHealthClub                        CategoryCode = "LIMITEDHEALTHCLUB"
	CategoryCodeDiving                                   CategoryCode = "DIVING"
	CategoryCodeWalkingTrack                             CategoryCode = "WALKING_TRACK"
	CategoryCodePaddleCourt                              CategoryCode = "PADDLE_COURT"
	CategoryCodeBoatTours                                CategoryCode = "BOAT_TOURS"
	CategoryCodeKidsGolfAcademy                          CategoryCode = "KIDS_GOLF_ACADEMY"
	CategoryCodeKidsBeachClub                            CategoryCode = "KIDS_BEACH_CLUB"
	CategoryCodeKidsEquestrianClub                       CategoryCode = "KIDS_EQUESTRIAN_CLUB"
	CategoryCodeLounge                                   CategoryCode = "LOUNGE"
)

type TransportMode string

// Enum values for TransportMode
const (
	Bicycle          TransportMode = "BICYCLE"
	Boat             TransportMode = "BOAT"
	Bus              TransportMode = "BUS"
	CableCar         TransportMode = "CABLE_CAR"
	CourtesyCar      TransportMode = "COURTESY_CAR"
	Car              TransportMode = "CAR"
	Carriage         TransportMode = "CARRIAGE"
	Helicopter       TransportMode = "HELICOPTER"
	Limousine        TransportMode = "LIMOUSINE"
	Metro            TransportMode = "METRO"
	Monorail         TransportMode = "MONORAIL"
	Motorbike        TransportMode = "MOTORBIKE"
	PackAnimal       TransportMode = "PACK_ANIMAL"
	Plane            TransportMode = "PLANE"
	Rickshaw         TransportMode = "RICKSHAW"
	Shuttle          TransportMode = "SHUTTLE"
	SedanChair       TransportMode = "SEDAN_CHAIR"
	Subway           TransportMode = "SUBWAY"
	Taxi             TransportMode = "TAXI"
	Train            TransportMode = "TRAIN"
	Walk             TransportMode = "WALK"
	WaterTaxi        TransportMode = "WATER_TAXI"
	OtherOrAlternate TransportMode = "OTHER_OR_ALTERNATE"
	ExpressTrain     TransportMode = "EXPRESS_TRAIN"
	Alternate        TransportMode = "ALTERNATE"
	Ferry            TransportMode = "FERRY"
)

type DistanceType string

// Enum values for DistanceType
const (
	Airways   DistanceType = "AIRWAYS"
	Roadways  DistanceType = "ROADWAYS"
	Railways  DistanceType = "RAILWAYS"
	Waterways DistanceType = "WATERWAYS"
	Birdseye  DistanceType = "BIRDSEYE"
)

type Segment string

// Enum values for Segment
const (
	SegmentAllSuite                   Segment = "ALL_SUITE"
	SegmentBudget                     Segment = "BUDGET"
	SegmentCorporateBusinessTransient Segment = "CORPORATE_BUSINESS_TRANSIENT"
	SegmentDeluxe                     Segment = "DELUXE"
	SegmentEconomy                    Segment = "ECONOMY"
	SegmentExtendedStay               Segment = "EXTENDED_STAY"
	SegmentFirstClass                 Segment = "FIRST_CLASS"
	SegmentLuxury                     Segment = "LUXURY"
	SegmentMeetingOrConvention        Segment = "MEETING_OR_CONVENTION"
	SegmentModerate                   Segment = "MODERATE"
	SegmentResidentialApartment       Segment = "RESIDENTIAL_APARTMENT"
	SegmentResort                     Segment = "RESORT"
	SegmentTourist                    Segment = "TOURIST"
	SegmentUpscale                    Segment = "UPSCALE"
	SegmentEfficiency                 Segment = "EFFICIENCY"
	SegmentStandard                   Segment = "STANDARD"
	SegmentMidscale                   Segment = "MIDSCALE"
	SegmentQuality                    Segment = "QUALITY"
	SegmentUnknown                    Segment = "UNKNOWN"
	SegmentMidscaleWithoutFAndB       Segment = "MIDSCALE_WITHOUT_F_AND_B"
	SegmentUpperUpscale               Segment = "UPPER_UPSCALE"
)

type HotelAreaType string

const (
	AreaAirport               HotelAreaType = "AIRPORT"
	AreaBeach                 HotelAreaType = "BEACH"
	AreaCity                  HotelAreaType = "CITY"
	AreaDowntown              HotelAreaType = "DOWNTOWN"
	AreaEast                  HotelAreaType = "EAST"
	AreaExpressway            HotelAreaType = "EXPRESSWAY"
	AreaLake                  HotelAreaType = "LAKE"
	AreaMountain              HotelAreaType = "MOUNTAIN"
	AreaNorth                 HotelAreaType = "NORTH"
	AreaResort                HotelAreaType = "RESORT"
	AreaRural                 HotelAreaType = "RURAL"
	AreaSouth                 HotelAreaType = "SOUTH"
	AreaSuburban              HotelAreaType = "SUBURBAN"
	AreaWest                  HotelAreaType = "WEST"
	AreaBeachfront            HotelAreaType = "BEACHFRONT"
	AreaOceanfront            HotelAreaType = "OCEANFRONT"
	AreaGulf                  HotelAreaType = "GULF"
	AreaBusinessDistrict      HotelAreaType = "BUSINESS_DISTRICT"
	AreaEntertainmentDistrict HotelAreaType = "ENTERTAINMENT_DISTRICT"
	AreaFinancialDistrict     HotelAreaType = "FINANCIAL_DISTRICT"
	AreaShoppingDistrict      HotelAreaType = "SHOPPING_DISTRICT"
	AreaTheatreDistrict       HotelAreaType = "THEATRE_DISTRICT"
	AreaCountryside           HotelAreaType = "COUNTRYSIDE"
	AreaBay                   HotelAreaType = "BAY"
	AreaMarina                HotelAreaType = "MARINA"
	AreaPark                  HotelAreaType = "PARK"
	AreaRiver                 HotelAreaType = "RIVER"
	AreaTouristSite           HotelAreaType = "TOURIST_SITE"
	AreaNorthSuburb           HotelAreaType = "NORTH_SUBURB"
	AreaSouthSuburb           HotelAreaType = "SOUTH_SUBURB"
	AreaEastSuburb            HotelAreaType = "EAST_SUBURB"
	AreaWestSuburb            HotelAreaType = "WEST_SUBURB"
	AreaWaterfront            HotelAreaType = "WATERFRONT"
	AreaSkiResort             HotelAreaType = "SKI_RESORT"
)

type HotelStatus string

const (
	StatusOpen                        HotelStatus = "OPEN"
	StatusClosed                      HotelStatus = "CLOSED"
	StatusPreOpening                  HotelStatus = "PRE_OPENING"
	StatusTest                        HotelStatus = "TEST"
	StatusPropertySuitableForChildren HotelStatus = "PROPERTY_SUITABLE_FOR_CHILDREN"
	StatusDeleted                     HotelStatus = "DELETED"
	StatusLocked                      HotelStatus = "LOCKED"
	StatusUnlocked                    HotelStatus = "UNLOCKED"
)

type (
	// * A point of interest is a specific location that someone may find useful or interesting that tourists visit, typically for its inherent or an exhibited natural or cultural value, historical significance, natural or built beauty, offering leisure and amusement.
	PointOfInterestResponse struct {
		Location            LocationResponse            `json:"location"`            // Indicates the location of the point of interest
		CategoryCode        CategoryCode                `json:"categoryCode"`        // Indicates the category to which the points of interest belongs to. It can contain values such as
		Description         string                      `json:"description"`         // Description of the point of interest
		Season              Period                      `json:"season"`              // Models a period of time between two dates and inclusive only of the days of the week specified.
		Contact             ContactResponse             `json:"contact"`             // A contact refers to the information that can be used to reach a person, a company or an organization.
		EligibilityForEntry []QualifiedFreeTextResponse `json:"eligibilityForEntry"` // Indicates the eligibility for entry to the point of interest
		OfficialWebsite     struct {
			Url string `json:"url"` // Indicates the URL of the website
		} `json:"officialWebsite"` // Indicates the official website of the point of interest
		OperatingHours   CalendarScheduleResponse  `json:"operatingHours"`   // As defined in: https://schema.org/Schedule A schedule defines a repeating time period used to describe a regularly occurring Event. At a minimum a schedule will specify repeatFrequency which describes the interval between occurences of the event. Additional information can be provided to specify the schedule more precisely. This includes identifying the day(s) of the week or month when the recurring event will take place, in addition to its start and end time. Schedules may also have start and end dates to indicate when they are active, e.g. to define a limited calendar of events.
		PriceEquation    []PricingEquationResponse `json:"priceEquation"`    // Indicates the price equation of the point of interest
		Transportations  []TransportationResponse  `json:"transportations"`  // Indicates the transportation info to reach the point of interest via various transportation modes
		LocationDistance LocationDistanceResponse  `json:"locationDistance"` // Indicates the location distance of the point of interest
		Media            []MediaResponse           `json:"media"`            // Indicates the media of the point of interest
		Hotel            HotelResponse             `json:"hotel"`            // Provides Information related to Hotel Calendar, Climate and Spoken Language.
		Basic            BasicResponse             `json:"basic"`            // By default, this Model would be returned in all successful cases. This information provides Information related to Hotel Name, Chain name and Hotel Id.
	}

	// BasicResponse struct {
	// 	Season struct {
	// 		OpenCalendar []Period `json:"openCalendar"` // Indicates the opening time and days of the property
	// 	} `json:"season"` // Indicates the season of the point of interest
	// 	HotelID                      string                         `json:"hotelId"`                      // Amadeus Property Code (8 chars). example: ADPAR001
	// 	ChainCode                    string                         `json:"chainCode"`                    // Brand (RT...) or Merchant (AD...)
	// 	BrandCode                    string                         `json:"brandCode"`                    // Brand (RT...) (Amadeus 2 chars Code). Small Properties distributed by Merchants may not have a Brand. Example - AD (Value Hotels) is the Provider/Merchant, and RT (Accor) is the Brand of the Property
	// 	DupeID                       string                         `json:"dupeId"`                       // Unique Property identifier of the physical hotel. One physical hotel can be represented by different Providers, each one having its own hotelID. This attribute allows a client application to group together hotels that are actually the same.
	// 	Name                         string                         `json:"name"`                         // Name of the point of interest
	// 	Rating                       string                         `json:"rating"`                       // Rating of the point of interest
	// 	Description                  QualifiedFreeTextResponse      `json:"description"`                  // Description of the point of interest
	// 	Amenities                    []AmenityResponse              `json:"amenities"`                    // Amenities of the point of interest
	// 	Media                        []MediaResponse                `json:"media"`                        // Media of the point of interest
	// 	DefaultSpokenLanguage        string                         `json:"defaultSpokenLanguage"`        // Describes the default language preferred or used at the property
	// 	ContextProvider              string                         `json:"contextProvider"`              // Describes the provider of the context of the point of interest
	// 	Contact                      []ContactResponse              `json:"contact"`                      // Contact of the point of interest
	// 	Location                     LocationResponse               `json:"location"`                     // Location of the point of interest
	// 	Altitude                     AltitudeResponse               `json:"altitude"`                     // From analytics, Metrics describe the exact numbers that make up the data
	// 	CategoryCode                 CategoryCode                   `json:"categoryCode"`                 // Category code of the point of interest
	// 	Segment                      Segment                        `json:"segment"`                      // Segment of the point of interest
	// 	Area                         []AreaResponse                 `json:"area"`                         // Geographical zone like City, Region, Country
	// 	ChainName                    string                         `json:"chainName"`                    // Name of the chain to which the hotel belongs to
	// 	BrandName                    string                         `json:"brandName"`                    // Name of the brand to which the hotel or hotel chain belongs to
	// 	Status                       HotelStatus                    `json:"status"`                       // Status of the hotel
	// 	HotelBusinessIdentifications BusinessIdentificationResponse `json:"hotelBusinessIdentifications"` // An business, can be idenfified via business identifiers, those business identifiers are defined by a body of authority thay could be local, national, transnational or supranational (like EU for the EU VAT number).
	// }

	// BusinessIdentificationResponse struct {
	// 	Identifiers []IdentifierResponse `json:"identifiers"` // Identifiers of the business
	// }

	// IdentifierResponse struct {
	// 	ID   string `json:"id"`   // Identifier id
	// 	Name string `json:"name"` // Identifier name
	// }

	// AreaResponse struct {
	// 	HotelAreaType HotelAreaType `json:"hotelAreaType"` // 'Indicates the category of the location. OTA Code Set LOC values are to be considered here. Can contain values such as
	// 	Name          string        `json:"name"`          // Label associated to the location (e.g. Eiffel Tower, Madison Square)
	// }

	// AltitudeResponse struct {
	// 	Unit  Unit `json:"unit"`  // Indicates the unit of the altitude
	// 	Value int  `json:"value"` // Indicates the value of the altitude
	// }

	// HotelResponse struct {
	// 	TaxID            string                     `json:"taxId"`        // Describes the unique tax identifier of a hotel property
	// 	CurrencyCode     []string                   `json:"currencyCode"` // Describes the currency code accepted at the property. Example : [EUR]
	// 	SpokenLanguages  []string                   // Describes the list of languages spoken at the property. Follows the standard of ISO 639-1 (Alpha-2). Example : [es]
	// 	TimeZone         TimeZoneResponse           `json:"timeZone"`         // Element defining a time zone
	// 	Climate          string                     `json:"climate"`          // Describes the climate at the location of the property. example: Dry
	// 	Certifications   []AwardsResponse           `json:"certifications"`   // Describes the certifications received by the Hotel
	// 	RelativeLocation []LocationDistanceResponse `json:"relativeLocation"` // To indicate the reference points from the hotel such as the distance to Airport, Bus Stations or Train Station.
	// 	Season           SeasonResponse             `json:"season"`           // Models a period of time between two dates and inclusive only of the days of the week specified.
	// 	Building         BuildingResponse           `json:"building"`         // Indicates the building of the hotel
	// }

	// BuildingResponse struct {
	// 	ArchitectureCode        ArchitectureCode `json:"architectureCode"`        // Denotes the architecture in which the property was built upon. Can contain values such as
	// 	BuildDate               string           `json:"buildDate"`               // Denotes the year at which the property was built. Format YYYY-MM-DD (ISO 8601)
	// 	RenovationDate          string           `json:"renovationDate"`          // Denotes the year at which the property was renovated. Format YYYY-MM-DD (ISO 8601)
	// 	NumberOfFloors          int              `json:"numberOfFloors"`          // Indicates the number of floors in the property
	// 	NumberOfRooms           int              `json:"numberOfRooms"`           // Indicates the number of rooms in the property
	// 	NumberOfExecutiveFloors int              `json:"numberOfExecutiveFloors"` // Indicates the number of Executive floors in the property
	// 	NumberOfBuildings       int              `json:"numberOfBuildings"`       // Indicates the number of buildings in the property
	// 	NumberOfElevators       int              `json:"numberOfElevators"`       // Indicates the number of elevators in the property
	// }

	// SeasonResponse struct {
	// 	ClosedSeasons   []Period `json:"closedSeasons"`   // Closed seasons of the hotel refers to the season where in the property is shut down
	// 	BlackoutSeasons []Period `json:"blackoutSeasons"` // Blackout dates of the hotel during which the hotel is open but no bookings are available
	// 	OpenCalendar    []Period `json:"openCalendar"`    // Indicates the opening time and days of the property
	// }

	// TimeZoneResponse struct {
	// 	ID                     string `json:"id"`                     // Unique id of the time zone. example: Europe/Paris
	// 	Name                   string `json:"name"`                   //Long name of the time zone. example: Central European Summer Time
	// 	Code                   string `json:"code"`                   // Time zone code. example: CEST
	// 	OffSet                 string `json:"offSet"`                 // Total offset from UTC including the Daylight Saving Time (DST) following ISO 8601 (https://en.wikipedia.org/wiki/ISO_8601) standard. example: +02:00
	// 	OffSetInSeconds        int    `json:"offSetInSeconds"`        // Total offset from UTC including the Daylight Saving Time (DST) in second. example: 7200
	// 	DstOffset              string `json:"dstOffset"`              // Indicates whether the day light savings is observed at the location. example: True
	// 	DstOffsetInSeconds     int    `json:"dstOffsetInSeconds"`     // Daylight Saving Time (DST) in second. 0 if the zone is not in the Daylight Saving time at specified date. example: -3600
	// 	ReferenceLocalDateTime string `json:"referenceLocalDateTime"` // Date and time used as reference to determine the time zone name, code, offset, and dstOffset following ISO 8601 (https://en.wikipedia.org/wiki/ISO_8601) standard. example: 2022-09-28T19:20:30
	// }

	// LocationDistanceResponse struct {
	// 	Destination LocationResponse   `json:"destination"` // Indicates the destination of the location distance
	// 	Distances   []DistanceResponse `json:"distances"`   // Indicates the distances from the point of interest to the destination
	// }

	// DistanceResponse struct {
	// 	Unit         Unit         `json:"unit"`         // Indicates the unit of the distance
	// 	Value        int          `json:"value"`        // Indicates the value of the distance
	// 	DistanceType DistanceType `json:"distanceType"` // Indicates the type of the distance
	// }

	// TransportationResponse struct {
	// 	TransportMode         TransportMode             `json:"transportMode"`
	// 	IsReservationRequired bool                      `json:"isReservationRequired"` // True if reservation is required in advance to board the transport
	// 	OperatingHours        CalendarScheduleResponse  `json:"operatingHours"`        // As defined in: https://schema.org/Schedule A schedule defines a repeating time period used to describe a regularly occurring Event. At a minimum a schedule will specify repeatFrequency which describes the interval between occurences of the event. Additional information can be provided to specify the schedule more precisely. This includes identifying the day(s) of the week or month when the recurring event will take place, in addition to its start and end time. Schedules may also have start and end dates to indicate when they are active, e.g. to define a limited calendar of events.
	// 	Description           string                    `json:"description"`           // Description of the transportation
	// 	PriceEquation         []PricingEquationResponse `json:"priceEquation"`         // Indicates the price equation of the transportation
	// 	Media                 []MediaResponse           `json:"media"`                 // Indicates the media of the transportation
	// }

	// PricingEquationResponse struct {
	// 	PricingMethod PricingMethod           `json:"pricingMethod"` // Indicates the pricing method of the point of interest
	// 	UnitPrice     ElementaryPriceResponse `json:"unitPrice"`     // Indicates the price of the point of interest per unit
	// }

	// ElementaryPriceResponse struct {
	// 	Amount              string           `json:"amount"`              // Indicates the amount of the price of the point of interest
	// 	Value               string           `json:"value"`               // Indicates the value of the price of the point of interest
	// 	DecimalPlaces       int              `json:"decimalPlaces"`       // Indicates the decimal places of the price of the point of interest
	// 	Currency            CurrencyResponse `json:"currency"`            // Indicates the currency of the price of the point of interest
	// 	ElementaryPriceType string           `json:"elementaryPriceType"` // Defines the type of price, eg. for base fare, total, grand total.
	// }

	// LocationResponse struct {
	// 	SubType  string                            `json:"subType"`  // Location sub-type (e.g. airport, port, rail-station, restaurant, atm...)
	// 	Name     string                            `json:"name"`     // Name of the location
	// 	IataCode string                            `json:"iataCode"` // IATA code of the location
	// 	GeoCode  sharedResponseDTO.GeoCodeResponse `json:"geoCode"`  // GeoCode of the location
	// }
	// Period struct {
	// 	Start *time.Time `json:"start"` // start date and time following ISO 8601 format
	// 	End   *time.Time `json:"end"`   // end date and time following ISO 8601 format
	// }
)
