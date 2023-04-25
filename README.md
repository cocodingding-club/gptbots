# GPTBOTS

## 구현 목표
- [ ] **0.실무나 독학시엔 제대로 써보기 힘든 고랭과 ChatGPT 유료 버전을 맘껏 가지고 놀기**
- 챗봇 생성 api
  - 챗봇 생성 API
- 도메인 추천
- [ ] 1.[GoDaddy](https://kr.godaddy.com/domainsearch/find?checkAvail=1&tmskey=%5B%40T%5Bsitecorecontent%3A%3Cfirstchild+id%3D%22%7B4745D3A8-B61D-4744-B2F6-72EBAB8565D7%7D%22+runpipeline%3D%22false%22+%2F%3E%5D%40T%5D&segment=repeat&domainToCheck=gabia.com)의 적합한 대안 제시처럼 검색어와 유사한 도메인 추천 기능을 ChatGPT를 통해 [가비아 도메인 검색화면](https://domain.gabia.com/regist/regist_step1.php)에 구현
  - 대강의 로직 구조
    - 검색어에 대해 ChatGPT로 유사 도메인 목록을 받는다
    - 고루틴을 이용해 병렬적으로 등록 여부를 확인한다(checkDomain).
    - 가장 먼저 등록 가능한 도메인에 대한 응답을 받는 즉시 모든 고루틴을 종료하고 응답 반환
  - 사용자 경험을 위해 어느 정도 성능(받아올 추천 목록의 길이, 등록 여부 응답 대기 시간)에 대한 최적화가 필요
  - API 호출 자체가 비용인 만큼 레디스를 이용한 캐싱도 검토
- [ ] 2.유사 도메인 추천을 개인화시키는 방향으로 고도화
  - 활용 가능한 데이터:
    - 사용자 아이디, 보유 도메인, 예약 도메인, 관심 도메인
  - 1번 항목에서 구현된 기능을 통해 위의 데이터에 대한 유사 도메인 목록 생성 후 추천
- [ ] 3.Go 어플리케이션의 배포 스크립트 및 사내서 보편적으로 이용할 만한 프로젝트 템플릿 구성
- [ ] 그 외 고려해봄직한 ML(?)스러운 접근법
  - ChatGPT를 이용한 서비스 사이트 컨텐츠 태깅 자동화 (Contents-based filtering)
    - Tag에 대한 CRUD API
    - 지정된 태그에 대해서만 태깅을 수행하도록 질의
  - 사용자의 아이디, 검색 이력, 보유 도메인 목록을 기반으로 키워드 추출(Keyword Extraction)
    - https://github.com/topics/keyword-extraction
    - 위에서 적당한 오픈소스 라이브러리를 이용해 추출하면 될듯
    - 추출한 키워드에 대해 1번 유사 도메인 추천 기능으로 추천 목록 추출
  - 사용자의 키워드 및 주된 관심사 태그에 대한 도메인 추천 (serendipity)
  
## 프로젝트 구조
```
- project
  - cmd [어플리케이션 진입점 / 빌드될 바이너리를 나눌 단위]
  - configs [설정]
    - db 접속 정보?
  - deploy [배포]
    - Dockerfile, Kubernetes 배포 스크립트 등
  - internal [외부에 노출시키지 않을 비즈니스 로직]
    - controller, service, repository 등등
    - REST API에 그치지 않고 graphql, grpc, websocket, mq 등등 다양하게 사용해보면 재밌을듯
  - pkg [외부 공개 라이브러리]
    - client
    - parser
  - vendor (?) [의존성 다운로드]
    - 오프라인 환경에서 코드 추적이 가능
  - web [프론트엔드]
    - api
    - [프론트엔드]
      - src
  - README.md
```
- 참고: https://github.com/golang-standards/project-layout/blob/master/README_ko.md
