func postHandler(w http.ResponseWriter, r *http.Request) {
    var data struct {
        FirstName string `json:"first_name"`
        LastName  string `json:"last_name"`
    }
    
    // Читаем тело запроса
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, "Ошибка при парсинге JSON", http.StatusBadRequest)
        return
    }
    
    // Выводим данные в консоль
    fmt.Println("Полученные данные:", data.FirstName, data.LastName)
    
    // Отправляем ответ
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Данные получены"))
}
