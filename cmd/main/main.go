package main

import (
	"DATA_FWD_TAP/config"
	"DATA_FWD_TAP/internal/app"
	"DATA_FWD_TAP/models"
	"DATA_FWD_TAP/models/structures"
	"fmt"
	"log"
	"os"

	"time"
)

func main() {

	args := os.Args
	serviceName := args[0]
	// Print the name of the program
	log.Println("Program :", args[0], " starts")

	/*configFilePath := filepath.Join("config", "config.ini")

	result := models.InitProcessSpace(configFilePath, "Server")
	if result != 0 {
		log.Fatalf("Failed to initialize process space: %d", result)
	}

	// fmt.Println(result)

	tokenValue := models.GetProcessSpaceValue("Host")
	if tokenValue == "" {
		log.Println("Token  not found")
	} else {
		fmt.Printf(" %s\n", tokenValue)
	}
	*/
	//----------------------------------------------

	//Here we are Loading  PostgreSQL configuration. this function "LoadPostgreSQLConfig" is inside the "PosgresqlConfig.go".
	cfg, err := config.LoadPostgreSQLConfig("C:/Users/devdu/go-workspace/data_fwd_tap/config/EnvConfig.ini")

	if err != nil {
		log.Fatalf("Failed to load PostgreSQL configuration: %v", err)
	}

	if config.GetDatabaseConnection(serviceName, *cfg) != 0 { // here we are establishing the connection with the database. so we are calling the function "config.GetDatabaseConnection(*cfg)" in "PosgresqlConfig.go".
		log.Fatalf("Failed to connect to database")
	}

	DB := config.GetDB()

	if DB == nil {
		log.Fatalf("Database connection is nil. Failed to connect to the database.")
	}

	app.Fn_bat_init(args[1:], DB) // here we are calling the function of cln_pack_clnt

	tm := &models.TransactionManager{}

	tranType := tm.FnBeginTran() // trasaction begins here. Here we are calling the "FnBeginTran()" of "Tuxlib.go".

	if tranType == -1 {
		log.Fatalf("Failed to begin transaction")
	}

	user := structures.User{
		Username: "new_user",
		Email:    "new_user@example.com",
	}
	if err := tm.Tran.Create(&user).Error; err != nil {
		log.Fatalf("Failed to insert new user: %v", err)
	}

	product := structures.Product{
		Name:          "New Product",
		Description:   "Description for New Product",
		Price:         39.99,
		StockQuantity: 200,
	}
	if err := tm.Tran.Create(&product).Error; err != nil {
		log.Fatalf("Failed to insert new product: %v", err)
	}

	order := structures.Order{
		UserID:    user.ID,
		OrderDate: time.Now(),
		Amount:    39.99,
		Status:    "Pending",
	}
	if err := tm.Tran.Create(&order).Error; err != nil {
		log.Fatalf("Failed to insert new order: %v", err)
	}

	orderItem := structures.OrderItem{
		OrderID:   order.ID,
		ProductID: product.ID,
		Quantity:  1,
		Price:     39.99,
	}
	if err := tm.Tran.Create(&orderItem).Error; err != nil {
		log.Fatalf("Failed to insert new order item: %v", err)
	}

	// Sample Read Operation
	var fetchedUser structures.User
	if err := tm.Tran.First(&fetchedUser, "username = ?", "new_user").Error; err != nil {
		log.Fatalf("Failed to fetch user: %v", err)
	}
	fmt.Printf("Fetched User: %+v\n", fetchedUser)

	// Sample Update Operation
	if err := tm.Tran.Model(&fetchedUser).Update("Email", "updated_user@example.com").Error; err != nil {
		log.Fatalf("Failed to update user email: %v", err)
	}

	// Sample Delete Operation
	if err := tm.Tran.Delete(&fetchedUser).Error; err != nil {
		log.Fatalf("Failed to delete user: %v", err)
	}

	if tm.FnCommitTran(tranType) != 0 {
		log.Fatalf("Failed to commit transaction")
	}

	fmt.Println("Transaction successfully committed")

}
