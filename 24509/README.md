# 상품 수여

## 문제 설명
- 각 과목별 성적이 가장 높은 학생에게 상품을 수여하는 문제

### 상품 수여 규칙
- 각 과목(국어, 영어, 수학, 과학)별로 가장 높은 점수를 받은 학생에게 수여
- 상품 수여 순서: 국어 → 영어 → 수학 → 과학
- 이미 상품을 받은 학생은 다른 과목의 수여 대상에서 제외
- 동점이면, 학생 번호가 낮은 사람에게 수여

## 입력
- N: 학생 수
- No: 학생 번호
- Scores[4]: 각 학생의 과목별 점수 배열 [국어, 영어, 수학, 과학]

## 출력
- 국어, 영어, 수학, 과학 순서로 상품을 받은 학생들의 번호

# 참고
- [백준 24509번](https://www.acmicpc.net/problem/24509)