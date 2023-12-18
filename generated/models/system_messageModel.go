package models

import "time"

type SystemMessage struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	SystemId string `json:"system_id"`
	MessageId string `json:"message_id"`
	HotelGroupId string `json:"hotel_group_id"`
	HotelId string `json:"hotel_id"`
	BusinessUnitId string `json:"business_unit_id"`
	RecruitmentSystemId string `json:"recruitment_system_id"`
	EmailTemplateId string `json:"email_template_id"`
	FormId string `json:"form_id"`
	JobPostingId string `json:"job_posting_id"`
	JobCandidateId string `json:"job_candidate_id"`
	JobAssessmentId string `json:"job_assessment_id"`
	JobInterviewId string `json:"job_interview_id"`
	JobOfferId string `json:"job_offer_id"`
	JobPlacementId string `json:"job_placement_id"`
	WorkRequestId string `json:"work_request_id"`
	WorkOrderId string `json:"work_order_id"`
  }