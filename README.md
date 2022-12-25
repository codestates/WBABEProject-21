# WBABEProject-21


## DB TABLE
![zzzz](https://user-images.githubusercontent.com/66409384/209439570-48e15c71-255d-42ea-9990-df93014cd166.png)


## 기능

### 피주문자 ("/api/vendor")
- 메뉴생성 ("/menu")[POST] <피주문자가 메뉴를 생성하는 기능>
- 메뉴수정 ("/update/menu/:id") <메뉴 수정 기능> 
- 메뉴삭제 ("/menu/:id")[DELETE] <메뉴 삭제 기능>
- Order 상태 업데이트 ("/update/status/:id") [PUT] <오더 상태 업데이트 기능>

### 주문자 ("/api/customer")
- 주문하기 ("/order") [POST] <주문기능, 3회 이상 주문을 하게되면 ORDER기록을 탐색해 5% 할인율 적용해 차감.>
- 메뉴조회 ("/menu/:page/:limit") [GET] <메뉴조회기능> 
- 주문내역조회 ("/order/:page/:limit/:userid") [GET] <오더기록조회기능>
- 주문메뉴변경 ("/order") [PUT] <주문메뉴수정, 단 Order의 state가 1이상일 경우 불가>
