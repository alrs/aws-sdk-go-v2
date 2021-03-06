// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CompareFacesInput struct {
	_ struct{} `type:"structure"`

	// A filter that specifies a quality bar for how much filtering is done to identify
	// faces. Filtered faces aren't compared. If you specify AUTO, Amazon Rekognition
	// chooses the quality bar. If you specify LOW, MEDIUM, or HIGH, filtering removes
	// all faces that don’t meet the chosen quality bar. The quality bar is based
	// on a variety of common use cases. Low-quality detections can occur for a
	// number of reasons. Some examples are an object that's misidentified as a
	// face, a face that's too blurry, or a face with a pose that's too extreme
	// to use. If you specify NONE, no filtering is performed. The default value
	// is NONE.
	//
	// To use quality filtering, the collection you are using must be associated
	// with version 3 of the face model or higher.
	QualityFilter QualityFilter `type:"string" enum:"true"`

	// The minimum level of confidence in the face matches that a match must meet
	// to be included in the FaceMatches array.
	SimilarityThreshold *float64 `type:"float"`

	// The input image as base64-encoded bytes or an S3 object. If you use the AWS
	// CLI to call Amazon Rekognition operations, passing base64-encoded image bytes
	// is not supported.
	//
	// If you are using an AWS SDK to call Amazon Rekognition, you might not need
	// to base64-encode image bytes passed using the Bytes field. For more information,
	// see Images in the Amazon Rekognition developer guide.
	//
	// SourceImage is a required field
	SourceImage *Image `type:"structure" required:"true"`

	// The target image as base64-encoded bytes or an S3 object. If you use the
	// AWS CLI to call Amazon Rekognition operations, passing base64-encoded image
	// bytes is not supported.
	//
	// If you are using an AWS SDK to call Amazon Rekognition, you might not need
	// to base64-encode image bytes passed using the Bytes field. For more information,
	// see Images in the Amazon Rekognition developer guide.
	//
	// TargetImage is a required field
	TargetImage *Image `type:"structure" required:"true"`
}

// String returns the string representation
func (s CompareFacesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CompareFacesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CompareFacesInput"}

	if s.SourceImage == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceImage"))
	}

	if s.TargetImage == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetImage"))
	}
	if s.SourceImage != nil {
		if err := s.SourceImage.Validate(); err != nil {
			invalidParams.AddNested("SourceImage", err.(aws.ErrInvalidParams))
		}
	}
	if s.TargetImage != nil {
		if err := s.TargetImage.Validate(); err != nil {
			invalidParams.AddNested("TargetImage", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CompareFacesOutput struct {
	_ struct{} `type:"structure"`

	// An array of faces in the target image that match the source image face. Each
	// CompareFacesMatch object provides the bounding box, the confidence level
	// that the bounding box contains a face, and the similarity score for the face
	// in the bounding box and the face in the source image.
	FaceMatches []CompareFacesMatch `type:"list"`

	// The face in the source image that was used for comparison.
	SourceImageFace *ComparedSourceImageFace `type:"structure"`

	// The value of SourceImageOrientationCorrection is always null.
	//
	// If the input image is in .jpeg format, it might contain exchangeable image
	// file format (Exif) metadata that includes the image's orientation. Amazon
	// Rekognition uses this orientation information to perform image correction.
	// The bounding box coordinates are translated to represent object locations
	// after the orientation information in the Exif metadata is used to correct
	// the image orientation. Images in .png format don't contain Exif metadata.
	//
	// Amazon Rekognition doesn’t perform image correction for images in .png
	// format and .jpeg images without orientation information in the image Exif
	// metadata. The bounding box coordinates aren't translated and represent the
	// object locations before the image is rotated.
	SourceImageOrientationCorrection OrientationCorrection `type:"string" enum:"true"`

	// The value of TargetImageOrientationCorrection is always null.
	//
	// If the input image is in .jpeg format, it might contain exchangeable image
	// file format (Exif) metadata that includes the image's orientation. Amazon
	// Rekognition uses this orientation information to perform image correction.
	// The bounding box coordinates are translated to represent object locations
	// after the orientation information in the Exif metadata is used to correct
	// the image orientation. Images in .png format don't contain Exif metadata.
	//
	// Amazon Rekognition doesn’t perform image correction for images in .png
	// format and .jpeg images without orientation information in the image Exif
	// metadata. The bounding box coordinates aren't translated and represent the
	// object locations before the image is rotated.
	TargetImageOrientationCorrection OrientationCorrection `type:"string" enum:"true"`

	// An array of faces in the target image that did not match the source image
	// face.
	UnmatchedFaces []ComparedFace `type:"list"`
}

// String returns the string representation
func (s CompareFacesOutput) String() string {
	return awsutil.Prettify(s)
}

const opCompareFaces = "CompareFaces"

// CompareFacesRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// Compares a face in the source input image with each of the 100 largest faces
// detected in the target input image.
//
// If the source image contains multiple faces, the service detects the largest
// face and compares it with each face detected in the target image.
//
// You pass the input and target images either as base64-encoded image bytes
// or as references to images in an Amazon S3 bucket. If you use the AWS CLI
// to call Amazon Rekognition operations, passing image bytes isn't supported.
// The image must be formatted as a PNG or JPEG file.
//
// In response, the operation returns an array of face matches ordered by similarity
// score in descending order. For each face match, the response provides a bounding
// box of the face, facial landmarks, pose details (pitch, role, and yaw), quality
// (brightness and sharpness), and confidence value (indicating the level of
// confidence that the bounding box contains a face). The response also provides
// a similarity score, which indicates how closely the faces match.
//
// By default, only faces with a similarity score of greater than or equal to
// 80% are returned in the response. You can change this value by specifying
// the SimilarityThreshold parameter.
//
// CompareFaces also returns an array of faces that don't match the source image.
// For each face, it returns a bounding box, confidence value, landmarks, pose
// details, and quality. The response also returns information about the face
// in the source image, including the bounding box of the face and confidence
// value.
//
// The QualityFilter input parameter allows you to filter out detected faces
// that don’t meet a required quality bar. The quality bar is based on a variety
// of common use cases. Use QualityFilter to set the quality bar by specifying
// LOW, MEDIUM, or HIGH. If you do not want to filter detected faces, specify
// NONE. The default value is NONE.
//
// To use quality filtering, you need a collection associated with version 3
// of the face model or higher. To get the version of the face model associated
// with a collection, call DescribeCollection.
//
// If the image doesn't contain Exif metadata, CompareFaces returns orientation
// information for the source and target images. Use these values to display
// the images with the correct image orientation.
//
// If no faces are detected in the source or target images, CompareFaces returns
// an InvalidParameterException error.
//
// This is a stateless API operation. That is, data returned by this operation
// doesn't persist.
//
// For an example, see Comparing Faces in Images in the Amazon Rekognition Developer
// Guide.
//
// This operation requires permissions to perform the rekognition:CompareFaces
// action.
//
//    // Example sending a request using CompareFacesRequest.
//    req := client.CompareFacesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CompareFacesRequest(input *CompareFacesInput) CompareFacesRequest {
	op := &aws.Operation{
		Name:       opCompareFaces,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CompareFacesInput{}
	}

	req := c.newRequest(op, input, &CompareFacesOutput{})

	return CompareFacesRequest{Request: req, Input: input, Copy: c.CompareFacesRequest}
}

// CompareFacesRequest is the request type for the
// CompareFaces API operation.
type CompareFacesRequest struct {
	*aws.Request
	Input *CompareFacesInput
	Copy  func(*CompareFacesInput) CompareFacesRequest
}

// Send marshals and sends the CompareFaces API request.
func (r CompareFacesRequest) Send(ctx context.Context) (*CompareFacesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CompareFacesResponse{
		CompareFacesOutput: r.Request.Data.(*CompareFacesOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CompareFacesResponse is the response type for the
// CompareFaces API operation.
type CompareFacesResponse struct {
	*CompareFacesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CompareFaces request.
func (r *CompareFacesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
