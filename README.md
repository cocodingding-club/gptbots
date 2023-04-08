# Gabia Recommender

## TODO
- [ ] [GoDaddy](https://kr.godaddy.com/domainsearch/find?checkAvail=1&tmskey=%5B%40T%5Bsitecorecontent%3A%3Cfirstchild+id%3D%22%7B4745D3A8-B61D-4744-B2F6-72EBAB8565D7%7D%22+runpipeline%3D%22false%22+%2F%3E%5D%40T%5D&segment=repeat&domainToCheck=gabia.com)의 적합한 대안 제시처럼 검색어와 유사한 도메인 추천 기능을 ChatGPT를 통해 [가비아 도메인 검색화면](https://domain.gabia.com/regist/regist_step1.php)에 구현
  - 간단하게 검색어에 대해 ChatGPT로 유사 도메인 목록을 받은 뒤 고루틴을 이용해 병렬적으로 등록 여부를 확인하면서(checkDomain) 가장 먼저 등록 가능한 도메인에 대한 응답을 받는 즉시 모든 고루틴을 종료하고 응답 반환
  - 사용자 경험을 위해 어느 정도 성능(받아올 추천 목록의 길이, 등록 여부 응답 대기 시간)을 조정하면서 최적화할 필요가 있어보임
  - API 호출 자체가 비용인 만큼 캐싱도 검토
- [ ] [가비아 도메인 메인 화면](https://domain.gabia.com/)을 개인화된 도메인 추천으로 고도화
  - 사용자의 아이디, 검색 이력, 보유 도메인 목록을 기반으로 키워드 추출(Keyword Extraction)
    - https://github.com/topics/keyword-extraction
    - 위에서 적당한 오픈소스 라이브러리를 이용해 추출하면 될듯
  - 추출한 키워드에 대해 앞서 구현한 유사 도메인 추천 기능으로 추천목록 추출
- [ ] Go 쿠버네티스 배포 스크립트 및 템플릿 작성

## Discussion
- [ ] 가비아의 다른 서비스와 시너지를 내거나 비슷한 추천 서비스를 제공할 수 있을까?
  - 예시 Usecase
    - 호스팅 혹은 클라우드 서비스를 이용 중인 고객의 관리콘솔 상에서 도메인 등록을 추천
    - 반대로 도메인 관리툴 상에서 호스팅, 클라우드 상품을 추천
  
## 프로젝트 구조
```
- project
  - cmd [어플리케이션]
  - configs [설정]
  - internal [비공개 비즈니스 로직]
  - pkg [외부 공개 라이브러리]
  - vendor [의존성 다운로드]
  - web [프론트엔드]
  - README.md
```
- 참고: https://github.com/golang-standards/project-layout/blob/master/README_ko.md
