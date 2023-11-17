// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package turnrpt

// tribeReport
//    tribeSection tribeScoutingReport tribeStatus tribeHumans tribeAnimals tribePossessions tribeSkills tribeMorale tribeWeight tribeSettlements
//
// tribeSection
//    tribeNumber COMMA COMMA tribeCurrentHex COMMA tribePreviousHex
//    tribeCurrentTurn COMMA tribeSeason COMMA tribeWeather
//    tribeNextTurn COMMA tribeNextTurnDate
//    tribePaymentReceived COMMA tribeTurnCost
//    tribeCredit
//    tribeGoodsTribe
//    tribeDesiredCommodities COMMENT
//
// tribeCredit
//    "Credit" COLON CURRENCY AMOUNT
// tribeCurrentHex
//    "Current Hex" EQUAL_SIGN hexNumber
// tribeCurrentTurn
//    "Current Turn" TURN_NUMBER OPEN_PAREN HASH NUMBER CLOSE_PAREN
// tribeDesiredCommodities
//    "Desired Commodities" COLON COMMENT
// tribeGoodsTribe
//    "Goods Tribe" COLON TEXT
// tribeNextTurn
//    "Next Turn" TURN_NUMBER OPEN_PAREN HASH NUMBER CLOSE_PAREN
// tribeNumber
//    "Tribe" TRIBE_NUMBER
// tribePaymentReceived
//    "Received" COLON CURRENCY AMOUNT
// tribePreviousHex
//    OPEN_PAREN "Previous Hex" EQUAL_SIGN hexNumber CLOSE_PAREN
// tribeScoutingReport
//    tribeScoutResults *
// tribeTurnCost
//    "Cost" COLON CURRENCY AMOUNT
// tribeSeason
//    "Winter"
// tribeWeather
//    "FINE"
//
// tribeScoutResults
//    "Scout" SCOUT_ID COLON "Scout" tribeScoutMovementResults (BACKSLASH tribeScoutMovementResults) *
// tribeScoutMovementResults
//    (scoutDirectionAndTerrain | scoutFailed) (COMMA somethingSomething) *
// scoutDirectionAndTerrain
//    DIRECTION DASH TERRAIN
// scoutFailed
//    somethingSomething
//
// tribeStatus
//    CLAN_NO "Status" COLON TERRAIN COMMA CLAN_NO
// tribeHumans
//    "Humans" humansCount +
// tribePossessions
//    tribeAnimals tribeMinerals tribeWarEquipment tribeFinishedGoods tribeRawMaterials tribeShips
// tribeAnimals
//    "Animals"        (animalItem NUMBER) +
// tribeMinerals
//    "Minerals"       (mineralItem NUMBER) +
// tribeWarEquipment
//    "War Equipment"  (warEquipmentItem NUMBER) +
// tribeFinishedGoods
//    "Finished Goods" (finishedGoodItem NUMBER) +
// tribeRawMaterials
//    "Raw Materials"  (rawMaterialItem NUMBER) +
// tribeShips
//    "None"
// tribeSkills
//    "Skills" COLON (skillCode NUMBER COMMA) +
// tribeMorale
//    "Morale" COLON NUMBER
// tribeWeight
//    "Weight" COLON NUMBER_WITH_COMMAS
// tribeSettlements
//     "Settlements" tribeSettlementHeading tribeSettlementRow *
// tribeSettlementHeading
//     "Hex Code" "Name" "Note" "Type" "Subtype"
// tribeSettlementRow
//     hexCode settlementName settlementNote settlementType settlementSubtype
//
// animalItem
//   "Cattle" | "Goat" | "Horse"
// finishedGoodItem
//   "Provs" | "Sling" | "Trap" | "Wagon"
// hexNumber
//     HASH_HASH HEX_NUMBER
// humansCount
//   ("Actives" | "Inactives" | "People" | "Warriors") NUMBER
// movementResults
//     DIRECTION "-" TERRAIN
//  or SPACES TEXT "," TEXT
// mineralItem
//    "Brass" | "Bronze" | "Coal" | "Iron" | "Silver"
// rawMaterialItem
//    "Bark" | "Bone" | "Gut" | "Leather" | "Log" | "Skin" | "Wax"
// skillCode
//    "Adm" | "BnW" | "Bon" | "Cur" | "Dip" | "Eco" | "Eng" | "For" | "Garr" | "Gut" | "Herd" | "Hunt" | "Int" | "Ldr" | "Ltr" | "Min" | "Qry" | "Sct" | "ShB" | "ShW" | "Skn" | "Tan" | "Wd"
// warEquipmentItem
//    "Club" | "Jerkin" | "Shield" | "Sword"
