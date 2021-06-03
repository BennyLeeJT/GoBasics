package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

type itemInfo struct {
	category int
	quantity int
	unitCost float64
}

// itemName := make(map[string]itemInfo) // kiv, check if this is the same way to create map
var itemName map[string]itemInfo

// var modifiedItemName map[string]itemInfo

func main() {
	// all variables declared here
	var (
		choiceMenu             int
		choiceGenerateReport   int
		totalHouseholdItemcost float64
		totalFoodItemcost      float64
		totalDrinksItemcost    float64

		// categorySliceSorted string

		newItemNameToAdd         string
		newItemCategoryNameToAdd string
		newItemCategoryNoToAdd   int
		newItemQuantityToAdd     int
		newItemUnitCostToAdd     float64

		loopMenu   bool = false
		selection       = [6]int{1: 6}
		categories      = [...]string{"Household", "Food", "Drinks"}
		itemsSlice      = []string{}
		eachItem        = []string{}

		existingItemNameToModify string
		itemNewName              string
		itemCategoryNameModifyTo string
		itemCategoryNoModifyTo   int
		itemQuantityModifyTo     int
		itemUnitCostModifyTo     float64

		existingItemCategoryName string
		existingItemCategoryNo   int
		existingItemQuantity     int
		existingUnitCost         float64

		noChangesSlice       = []string{}
		noChangeItemName     string
		noChangeCategoryName string
		noChangeQuantity     string
		noChangeUnitCost     string

		existingItemNameToDelete string

		maintainItemNameInfo bool = false
	)

	// fmt.Println("initial choice declared defaulted to value : ", choiceMenu) // check default to 0
	// fmt.Println("initial loopMenu declared defaulted to value : ", loopMenu) // check default to false
	// fmt.Println("slice categories contains:", categories) // check. i need to use declared slice to rid of error
	fmt.Println("")

	for {
		// fmt.Println("first line after infinite for loop") // check

		if choiceMenu == 0 || loopMenu == true {

			fmt.Println("Shopping List Application\n" +
				"=========================\n" +
				"1. View entire shopping list.\n" +
				"2. Generate Shopping List Report.\n" +
				"3. Add Items.\n" +
				"4. Modify Items.\n" +
				"5. Delete Item.\n" +
				"6. Print Current Data.\n" +
				"Select your choice by indicating the number :")

			fmt.Scanln(&choiceMenu)
			fmt.Printf("** You have chosen Menu choice no. %v **", choiceMenu) // check
			fmt.Println("")

			if maintainItemNameInfo == false { // declared as false initially
				// bool condition check added after delete feature done to capture empty map.

				// check if program is first time running, if yes, create new map, else don't recreate
				// what has been added previously will then appear if user goes back view shopping list again
				if len(itemName) == 0 {
					itemName = make(map[string]itemInfo)

					itemName["Fork"] = itemInfo{0, 4, 3}
					itemName["Plates"] = itemInfo{0, 4, 3}
					itemName["Cups"] = itemInfo{0, 5, 3}
					itemName["Bread"] = itemInfo{1, 2, 2}
					itemName["Cake"] = itemInfo{1, 3, 1}
					itemName["Coke"] = itemInfo{2, 5, 2}
					itemName["Sprite"] = itemInfo{2, 5, 2}
				}
			}

			// validate choice of user, to go back to shopping list menu if invalid
			if choiceMenu > len(selection) || choiceMenu < 1 {
				fmt.Println("**>> Note : Please select option 1 to 6 only <<**")
				fmt.Println("")
			} else {

				if choiceMenu == 1 {
					for key, value := range itemName {
						if value.category == 0 {
							fmt.Printf("Category: Household - Item: %v Quantity: %v Unit Cost: %v\n", key, value.quantity, value.unitCost)
						} else if value.category == 1 {
							fmt.Printf("Category: Food - Item: %v Quantity: %v Unit Cost: %v\n", key, value.quantity, value.unitCost)
						} else if value.category == 2 {
							fmt.Printf("Category: Drinks - Item: %v Quantity: %v Unit Cost: %v\n", key, value.quantity, value.unitCost)
						}

						// fmt.Println(itemName["Fork"].category) // test accessing value of stuct, as a value in a map by searching the key
						// fmt.Printf("Category: %v - Item: %v\n", key, value) // test
					}
					fmt.Println("")

					// END OF CHOICE MENU 1 DISPLAY SHOPPING LIST
				} else if choiceMenu == 2 { // Generate Shopping List Report
					fmt.Println("Generate Report\n" +
						"1. Total Cost of each category.\n" +
						"2. List of item by category.\n" +
						"3. Main Menu.\n" +
						"\n" +
						"Choose your report by indicating the number :")

					fmt.Scanln(&choiceGenerateReport)
					fmt.Printf("** You have chosen Generate Report choice no. %v **", choiceGenerateReport) // check
					fmt.Println("")
					// fmt.Println("itemName: ", itemName)

					if choiceGenerateReport == 1 {
						// fmt.Println("calculation time: ", choiceGenerateReport) // check

						for _, value := range itemName {
							// fmt.Println("current value.category: ", value.category) // check
							if value.category == 0 {
								householdItemCost := float64(value.quantity) * value.unitCost
								totalHouseholdItemcost += householdItemCost
								// fmt.Println("key: ", key)                                       // check
								// fmt.Println("itemCost: ", householdItemCost)                    // check
								// fmt.Println("totalHouseholdItemcost: ", totalHouseholdItemcost) // check
							} else if value.category == 1 {
								foodItemCost := float64(value.quantity) * value.unitCost
								totalFoodItemcost += foodItemCost
								// fmt.Println("key: ", key) // check
								// fmt.Println("itemCost: ", foodItemCost) // check
								// fmt.Println("totalHouseholdItemcost: ", totalFoodItemcost) // check
							} else if value.category == 2 {
								drinksItemCost := float64(value.quantity) * value.unitCost
								totalDrinksItemcost += drinksItemCost
								// fmt.Println("key: ", key) // check
								// fmt.Println("itemCost: ", drinksItemCost) // check
								// fmt.Println("totalHouseholdItemcost: ", totalDrinksItemcost) // check
							}
						}

						fmt.Println("Total cost by Category.")
						fmt.Printf("1. Household cost : %v\n", totalHouseholdItemcost)
						fmt.Printf("2. Food cost : %v\n", totalFoodItemcost)
						fmt.Printf("3. Drinks cost : %v\n", totalDrinksItemcost)
						fmt.Println("")

						totalHouseholdItemcost = 0 // need to reset total if program reruns from top
						totalFoodItemcost = 0      // need to reset total if program reruns from top
						totalDrinksItemcost = 0    // need to reset total if program reruns from top

					} else if choiceGenerateReport == 2 { // List of item by category
						fmt.Println("List by Category.")
						for key, value := range itemName {

							if value.category == 0 {
								// fmt.Println("if started")
								eachItem = []string{
									"Category: " +
										categories[0] +
										" - Item: " +
										key +
										" Quantity: " +
										strconv.Itoa(value.quantity) +
										" Unit Cost: " +
										fmt.Sprintf("%.2f", math.Round(value.unitCost*100)/100),
								}
								// fmt.Println("eachItem : ", eachItem)
								// fmt.Println("eachItem[0] : ", eachItem[0])
								// fmt.Println("itemsSlice : ", itemsSlice)
								// itemsSlice = [...]string{"Household", "Food", "Drinks"}

								// fmt.Printf("Category: Household - Item: %v Quantity: %v Unit Cost: %v\n", key, value.quantity, value.unitCost)
							}

							if value.category == 1 {
								// fmt.Println("if started")
								eachItem = []string{
									"Category: " +
										categories[1] +
										" - Item: " +
										key +
										" Quantity: " +
										strconv.Itoa(value.quantity) +
										" Unit Cost: " +
										fmt.Sprintf("%.2f", math.Round(value.unitCost*100)/100),
								}
								// fmt.Println("eachItem : ", eachItem)
								// fmt.Println("eachItem[0] : ", eachItem[0])

								// itemsSlice = itemSlice + eachItem
								// itemsSlice = itemsSlice + eachItem[0]
								// itemsSlice = append(itemsSlice, eachItem[0])
								// fmt.Printf("Category: Food - Item: %v Quantity: %v Unit Cost: %v\n", key, value.quantity, value.unitCost)
							}

							if value.category == 2 {
								// fmt.Println("if started")
								eachItem = []string{
									"Category: " +
										categories[2] +
										" - Item: " +
										key +
										" Quantity: " +
										strconv.Itoa(value.quantity) +
										" Unit Cost: " +
										fmt.Sprintf("%.2f", math.Round(value.unitCost*100)/100),
								}
								// fmt.Println("eachItem : ", eachItem)
								// fmt.Println("eachItem[0] : ", eachItem[0])

								// itemsSlice = itemSlice + eachItem
								// itemsSlice = itemsSlice + eachItem[0]
								// itemsSlice = append(itemsSlice, eachItem[0])
								// fmt.Printf("Category: Drinks - Item: %v Quantity: %v Unit Cost: %v\n", key, value.quantity, value.unitCost)
							}
							itemsSlice = append(itemsSlice, eachItem[0])

						}

						sort.Strings(itemsSlice)
						for i := range itemsSlice {
							fmt.Printf("%v\n", itemsSlice[i])
						}

						fmt.Println("")
						itemsSlice = nil

					}
					// pressing enter or anything not 1 and 2 will go back anyways
					// else if choiceGenerateReport == 3 {
					// 	break
					// goes back to infinite for loop
					// }

					// END OF CHOICE MENU 2 GENERATE REPORT with 3 options
				} else if choiceMenu == 3 { // Add Items to existing categories
					fmt.Println("What is the name of your item?")
					fmt.Scanln(&newItemNameToAdd)

					fmt.Println("What existing category does it belong to?")
					fmt.Scanln(&newItemCategoryNameToAdd)

					if newItemCategoryNameToAdd == "Household" {
						newItemCategoryNoToAdd = 0
					} else if newItemCategoryNameToAdd == "Food" {
						newItemCategoryNoToAdd = 1
					} else if newItemCategoryNameToAdd == "Drinks" {
						newItemCategoryNoToAdd = 2
					}

					fmt.Println("How many units are there?")
					fmt.Scanln(&newItemQuantityToAdd)

					fmt.Println("How much does it cost per unit")
					fmt.Scanln(&newItemUnitCostToAdd)

					fmt.Println("Adding New Item... please wait")
					fmt.Println("newItemNameToAdd : ", newItemNameToAdd)
					fmt.Println("newItemCategoryNameToAdd : ", newItemCategoryNameToAdd)
					fmt.Println("newItemCategoryNoToAdd : ", newItemCategoryNoToAdd)
					fmt.Println("newItemQuantityToAdd : ", newItemQuantityToAdd)
					fmt.Println("newItemUnitCostToAdd : ", newItemUnitCostToAdd)

					itemName[newItemNameToAdd] = itemInfo{
						newItemCategoryNoToAdd,
						newItemQuantityToAdd,
						newItemUnitCostToAdd,
					}

					fmt.Println("New Item successfully added to Map!")
					fmt.Println("")
					fmt.Println("itemName map updated as follows: \n", itemName) // check item is inside map
					fmt.Println("")

					// END OF CHOICE MENU 3 ADD NEW ITEMS
				} else if choiceMenu == 4 { // Modify Items
					fmt.Println("Modify Items.")
					fmt.Println("Which item would you wish to modify?")
					fmt.Scanln(&existingItemNameToModify)

					// fmt.Println("itemName[existingItemNameToModify] : ", itemName[existingItemNameToModify]) // check able to access and retrieve

					_, ok := itemName[existingItemNameToModify]
					// fmt.Println("checkExist : ", _) // check struct info
					// fmt.Println("ok : ", ok) // check bool

					// fmt.Printf("test print %v\n", itemName[existingItemNameToModify].category) // check category retrieve
					fmt.Println("")

					existingItemCategoryName = categories[itemName[existingItemNameToModify].category]
					existingItemCategoryNo = itemName[existingItemNameToModify].category
					existingItemQuantity = itemName[existingItemNameToModify].quantity
					existingUnitCost = math.Round(itemName[existingItemNameToModify].unitCost)
					// if ok == true {
					// 	fmt.Println("Current item name is " +
					// 		existingItemNameToModify +
					// 		" - Category is " +
					// 		categories[itemName[existingItemNameToModify].category] +
					// 		" - Quantity is " +
					// 		strconv.Itoa(itemName[existingItemNameToModify].quantity) +
					// 		" - Unit Cost " +
					// 		fmt.Sprintf("%.2f", math.Round(itemName[existingItemNameToModify].unitCost)),
					// 	)
					// 	fmt.Println("")
					// }

					if ok == true {
						fmt.Println("Current item name is " +
							existingItemNameToModify +
							" - Category is " +
							existingItemCategoryName +
							" - Quantity is " +
							strconv.Itoa(existingItemQuantity) +
							" - Unit Cost " +
							fmt.Sprintf("%.2f", existingUnitCost),
						)
						fmt.Println("")
					}

					fmt.Println("Enter new Name. Enter blank for no change.")
					fmt.Scanln(&itemNewName)
					// fmt.Println("itemNewName : ", itemNewName) // check
					// fmt.Println("length of itemNewName : ", len(itemNewName)) // check

					if len(itemNewName) == 0 {
						noChangeItemName = "No changes to item name made"
						noChangesSlice = append(noChangesSlice, noChangeItemName)
						itemNewName = existingItemNameToModify
					}
					// fmt.Println("itemNewName : ", itemNewName) // check

					fmt.Println("Enter new Category name from existing categories. Enter blank for no change.")
					fmt.Scanln(&itemCategoryNameModifyTo)
					// fmt.Println("length of itemCategoryNameModifyTo : ", len(itemCategoryNameModifyTo)) // check
					// fmt.Println("existingItemCategoryName : ", existingItemCategoryName) // check
					// itemCategoryNameModifyTo = existingItemCategoryName
					// fmt.Println("itemCategoryNameModifyTo : ", itemCategoryNameModifyTo) // check

					// fmt.Printf("%T, %v\n", itemCategoryNameModifyTo, itemCategoryNameModifyTo) // check
					// if bool(itemCategoryNameModifyTo) == false {

					// }
					if len(itemCategoryNameModifyTo) == 0 {
						noChangeCategoryName = "No changes to category made"
						noChangesSlice = append(noChangesSlice, noChangeCategoryName)
						// fmt.Println("len(itemCategoryNameModifyTo) : ZERO") // check
						itemCategoryNameModifyTo = existingItemCategoryName
						itemCategoryNoModifyTo = existingItemCategoryNo
					} else {
						if itemCategoryNameModifyTo == "Household" {
							// var itemCategoryNoModifyTo int
							// itemCategoryNoModifyTo = 1
							itemCategoryNoModifyTo = 0
						} else if itemCategoryNameModifyTo == "Food" {
							itemCategoryNoModifyTo = 1
						} else if itemCategoryNameModifyTo == "Drinks" {
							itemCategoryNoModifyTo = 2
						}
					}

					// fmt.Println("itemCategoryNoModifyTo : ", itemCategoryNoModifyTo) // need this initially so that var is used

					fmt.Println("Enter new Quantity. Enter blank for no change.")
					// fmt.Scanln(&itemQuantityModifyTo) // check
					fmt.Println("itemQuantityModifyTo : ", itemQuantityModifyTo)
					if itemQuantityModifyTo == 0 { // cannot use len here for int type
						// itemQuantityModifyTo = existingItemQuantity
						noChangeQuantity = "No changes to quantity made"
						noChangesSlice = append(noChangesSlice, noChangeQuantity)
						itemQuantityModifyTo = itemName[existingItemNameToModify].quantity
					}

					fmt.Println("Enter new Unit Cost. Enter blank for no change.")
					fmt.Scanln(&itemUnitCostModifyTo)
					// fmt.Println("itemUnitCostModifyTo : ", itemUnitCostModifyTo) // check
					if itemUnitCostModifyTo == 0.00 { // cannot use len here for int type
						// itemQuantityModifyTo = existingUnitCost
						noChangeUnitCost = "No changes to unit cost made"
						noChangesSlice = append(noChangesSlice, noChangeUnitCost)
						itemUnitCostModifyTo = existingUnitCost
					}

					fmt.Println("noChangesSlice : ", noChangesSlice)
					for i := 0; i < len(noChangesSlice); i++ {
						fmt.Println(noChangesSlice[i])

					}
					fmt.Println("")

					delete(itemName, existingItemNameToModify)
					// fmt.Println("itemName after delete : ", itemName) // check
					itemName[itemNewName] = itemInfo{itemCategoryNoModifyTo, itemQuantityModifyTo, itemUnitCostModifyTo}
					// fmt.Println("itemName after added modified data : ", itemName) // check
					// fmt.Println("modifieditemName : ", modifieditemName)

					// END OF CHOICE MENU 4 MODIFY ITEMS
				} else if choiceMenu == 5 { // Delete Items
					fmt.Println("** Delete Item **")
					fmt.Println("Enter item to delete?")
					fmt.Scanln(&existingItemNameToDelete)

					// for key, value := range itemName {
					// 	if key ==
					// 	fmt.Println("value : ", value)
					// }

					if _, exists := itemName[existingItemNameToDelete]; exists {
						// fmt.Println("checkKeyitemInfo :", _)
						fmt.Println("existingItemNameToDelete :", existingItemNameToDelete)
						fmt.Println("exists :", exists) // check bool return
						delete(itemName, existingItemNameToDelete)
						fmt.Printf("Deleted %v", existingItemNameToDelete)
						fmt.Println("")
					} else {
						fmt.Println("Item not found. Nothing to delete!")
					}
					fmt.Println("")

					// END OF CHOICE MENU 5 DELETE ITEMS
				} else if choiceMenu == 6 { // Print current data fields
					fmt.Println("** Print Current Data **")
					fmt.Println("len(itemName) : ", len(itemName))
					if len(itemName) == 0 {
						fmt.Println("No data found!")
					} else {
						for key, value := range itemName {
							// if len(itemName) == 0 {
							// 	fmt.Println("No data found!")
							// } else {
							fmt.Printf("%v - %v\n", key, value)
							// }

						}
					}

					maintainItemNameInfo = true
					fmt.Println("")
				} // END OF CHOICE MENU 6 PRINT CURRENT DATA
			}

		} // end first if check

		loopMenu = true
		// fmt.Println("loopMenu is now : ", loopMenu) // check
		// goes back to infinite for loop
	}
}
